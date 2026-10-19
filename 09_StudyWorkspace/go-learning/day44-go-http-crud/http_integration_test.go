package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestMux 使用现有的业务 Handler；没有重新实现 CRUD。
// 当前测试路由与 main.go 保持一致；未来可抽取生产与测试共用的 newRouter(repo)。
func newTestMux(repo *StudentRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /students", studentsHandler(repo))
	mux.HandleFunc("GET /students/{id}", studentByIDHandler(repo))
	mux.HandleFunc("POST /students", createStudentHandler(repo))
	mux.HandleFunc("PUT /students/{id}", updateStudentHandler(repo))
	mux.HandleFunc("DELETE /students/{id}", deleteStudentHandler(repo))
	return mux
}

// performRequest 在内存中执行一条 HTTP 请求；不会开启监听端口。
func performRequest(t *testing.T, handler http.Handler, method, path, jsonBody string) (int, http.Header, []byte) {
	t.Helper()

	var body io.Reader
	if jsonBody != "" {
		body = strings.NewReader(jsonBody)
	}
	req := httptest.NewRequest(method, path, body)
	if jsonBody != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	resp := recorder.Result()
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read HTTP response: %v", err)
	}
	return resp.StatusCode, resp.Header, responseBody
}

func requireStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("HTTP status: want %d, got %d", want, got)
	}
}

func requireStudentJSON(t *testing.T, body []byte, want Student) {
	t.Helper()
	var got Student
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode student JSON (%q): %v", string(body), err)
	}
	if got != want {
		t.Errorf("student JSON: want %+v, got %+v", want, got)
	}
}

func requireErrorJSON(t *testing.T, body []byte, want string) {
	t.Helper()
	var got ErrorResponse
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode error JSON (%q): %v", string(body), err)
	}
	if got.Error != want {
		t.Errorf("error JSON: want %q, got %q", want, got.Error)
	}
}

func TestGetStudentHTTPWithSQLite(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	want := Student{ID: 9001, Name: "Eva", Age: 23}
	if err := repo.Create(want); err != nil {
		t.Fatalf("seed student: %v", err)
	}

	status, headers, body := performRequest(t, newTestMux(repo), http.MethodGet, "/students/9001", "")
	requireStatus(t, status, http.StatusOK)
	if got := headers.Get("Content-Type"); !strings.HasPrefix(got, "application/json") {
		t.Errorf("Content-Type: want application/json, got %q", got)
	}
	requireStudentJSON(t, body, want)
}

func TestListStudentsHTTPWithSQLite(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	mux := newTestMux(repo)

	status, _, body := performRequest(t, mux, http.MethodGet, "/students", "")
	requireStatus(t, status, http.StatusOK)
	if strings.TrimSpace(string(body)) != "[]" {
		t.Errorf("empty list: want [], got %q", string(body))
	}

	want := Student{ID: 9005, Name: "Lin", Age: 24}
	if err := repo.Create(want); err != nil {
		t.Fatalf("seed student: %v", err)
	}
	status, _, body = performRequest(t, mux, http.MethodGet, "/students", "")
	requireStatus(t, status, http.StatusOK)

	var got []Student
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode student list: %v", err)
	}
	if len(got) != 1 || got[0] != want {
		t.Errorf("list: want [%+v], got %+v", want, got)
	}
}

func TestCreateStudentHTTPWithSQLite(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	mux := newTestMux(repo)
	want := Student{ID: 9002, Name: "Alice", Age: 22}

	status, headers, body := performRequest(t, mux, http.MethodPost, "/students", `{"id":9002,"name":"Alice","age":22}`)
	requireStatus(t, status, http.StatusCreated)
	if got := headers.Get("Location"); got != "/students/9002" {
		t.Errorf("Location: want %q, got %q", "/students/9002", got)
	}
	requireStudentJSON(t, body, want)

	saved, err := repo.GetByID(9002)
	if err != nil {
		t.Fatalf("student not persisted: %v", err)
	}
	if saved != want {
		t.Errorf("database: want %+v, got %+v", want, saved)
	}
}

func TestUpdateStudentHTTPWithSQLite(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	before := Student{ID: 9003, Name: "Eva", Age: 22}
	if err := repo.Create(before); err != nil {
		t.Fatalf("seed student: %v", err)
	}

	// 重要：当前 PUT Handler 使用 DisallowUnknownFields，
	// 请求体只能包含 name 和 age；ID 由 URL /students/9003 提供。
	want := Student{ID: 9003, Name: "Eva", Age: 23}
	status, _, body := performRequest(t, newTestMux(repo), http.MethodPut, "/students/9003", `{"name":"Eva","age":23}`)
	requireStatus(t, status, http.StatusOK)
	requireStudentJSON(t, body, want)

	saved, err := repo.GetByID(9003)
	if err != nil {
		t.Fatalf("get updated student: %v", err)
	}
	if saved != want {
		t.Errorf("database: want %+v, got %+v", want, saved)
	}
}

func TestDeleteStudentHTTPWithSQLite(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	student := Student{ID: 9004, Name: "David", Age: 24}
	if err := repo.Create(student); err != nil {
		t.Fatalf("seed student: %v", err)
	}
	mux := newTestMux(repo)

	status, _, body := performRequest(t, mux, http.MethodDelete, "/students/9004", "")
	requireStatus(t, status, http.StatusNoContent)
	if len(body) != 0 {
		t.Errorf("204 response body should be empty; got %q", string(body))
	}

	_, err := repo.GetByID(9004)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("after DELETE: want sql.ErrNoRows, got %v", err)
	}

	status, _, body = performRequest(t, mux, http.MethodDelete, "/students/9004", "")
	requireStatus(t, status, http.StatusNotFound)
	requireErrorJSON(t, body, "student not found")
}

func TestHTTPValidationAndMissingResources(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	mux := newTestMux(repo)

	cases := []struct {
		name, method, path, jsonBody string
		wantStatus                   int
		wantError                    string
	}{
		{"GET invalid ID", http.MethodGet, "/students/abc", "", 400, "invalid student id"},
		{"GET missing", http.MethodGet, "/students/9999", "", 404, "student not found"},
		{"POST invalid data", http.MethodPost, "/students", `{"id":0,"name":"Alice","age":22}`, 400, "invalid student data"},
		{"POST invalid JSON", http.MethodPost, "/students", `{"id":`, 400, "invalid JSON body"},
		{"PUT invalid ID", http.MethodPut, "/students/abc", `{"name":"Eva","age":23}`, 400, "invalid student id"},
		{"PUT missing", http.MethodPut, "/students/9999", `{"name":"Eva","age":23}`, 404, "student not found"},
		{"PUT unknown id field", http.MethodPut, "/students/9003", `{"id":9003,"name":"Eva","age":23}`, 400, "invalid JSON body"},
		{"DELETE invalid ID", http.MethodDelete, "/students/abc", "", 400, "invalid student id"},
		{"DELETE missing", http.MethodDelete, "/students/9999", "", 404, "student not found"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			status, _, body := performRequest(t, mux, tc.method, tc.path, tc.jsonBody)
			requireStatus(t, status, tc.wantStatus)
			requireErrorJSON(t, body, tc.wantError)
		})
	}
}

func TestCreateStudentConflict(t *testing.T) {
	repo := newIsolatedStudentRepo(t)
	want := Student{ID: 9006, Name: "Alice", Age: 22}
	if err := repo.Create(want); err != nil {
		t.Fatalf("seed student: %v", err)
	}

	status, _, body := performRequest(t, newTestMux(repo), http.MethodPost, "/students", `{"id":9006,"name":"Alice","age":22}`)
	requireStatus(t, status, http.StatusConflict)
	requireErrorJSON(t, body, "student id already exists")

	saved, err := repo.GetByID(9006)
	if err != nil || saved != want {
		t.Fatalf("original record unexpectedly modified: saved=%+v, err=%v", saved, err)
	}
}
