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

func TestDay46InvalidStudentID(t *testing.T) {
	repo := newIsolatedStudentRepo(t)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students/{id}",
		studentByIDHandler(repo))
	// 1. 自己创建 GET /students/abc 请求
	req := httptest.NewRequest(
		http.MethodGet,
		"/students/abc",
		nil,
	)
	// 2. 自己创建 ResponseRecorder
	recorder := httptest.NewRecorder()
	// 3. 调用 mux.ServeHTTP()
	mux.ServeHTTP(recorder, req)
	// 4. 获取响应并检查状态码是否为 400
	result := recorder.Result().StatusCode
	if result != 400 {
		t.Fatalf("required status code 400, but got %v", result)
	}
}

//标准写法
// func TestDay46InvalidStudentID(t *testing.T) {
//     repo := newIsolatedStudentRepo(t)

//     mux := http.NewServeMux()
//     mux.HandleFunc(
//         "GET /students/{id}",
//         studentByIDHandler(repo),
//     )

//     // 构造非法 ID 请求
//     req := httptest.NewRequest(
//         http.MethodGet,
//         "/students/abc",
//         nil,
//     )

//     recorder := httptest.NewRecorder()

//     // 执行路由和 Handler
//     mux.ServeHTTP(recorder, req)

//     // 获取 HTTP 响应
//     resp := recorder.Result()
//     defer resp.Body.Close()

//	    // 验证 400
//	    if resp.StatusCode != http.StatusBadRequest {
//	        t.Fatalf(
//	            "expected status 400, got %d",
//	            resp.StatusCode,
//	        )
//	    }
//	}
func TestDay47GetStudentWithSQLite(t *testing.T) {
	// TODO 1：创建独立测试数据库 Repository
	repo := newIsolatedStudentRepo(t)
	// TODO 2：使用 Repository 向数据库插入学生
	// ID=1001, Name="Alice", Age=20
	// repo.Create(Student{ID: 1001, Name: "Alice", Age: 20})
	err := repo.Create(Student{
		ID:   1001,
		Name: "Alice",
		Age:  20,
	})
	if err != nil {
		t.Fatalf("create student failed: %v", err)
	}
	// TODO 3：创建 ServeMux，注册 GET /students/{id}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /students/{id}", studentByIDHandler(repo))
	// TODO 4：构造 GET /students/1001 请求
	req := httptest.NewRequest(http.MethodGet, "/students/1001", nil)
	// TODO 5：创建 Recorder 并执行 ServeHTTP
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)
	// TODO 6：检查 HTTP 状态码为 200
	result := recorder.Result()
	if result.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 but got %d", result.StatusCode)
	}
	// TODO 7：解析响应 JSON，检查学生字段
	// want := Student{ID: 1001, Name: "Alice", Age: 20}
	// got := Student{}
	// if json.NewDecoder(result.Body).Decode(&got); got != want {
	// 	t.Fatalf("expected 1001 Alice 20 but got %v", got)
	// }
	var got Student

	err = json.NewDecoder(result.Body).Decode(&got)
	if err != nil {
		t.Fatalf("decode response failed: %v", err)
	}

	want := Student{ID: 1001, Name: "Alice", Age: 20}

	if got != want {
		t.Fatalf("expected %+v, got %+v", want, got)
	}
}
