package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试 GET /students/1
func TestGetStudentOK(t *testing.T) {
	handler := newStudentHandler()

	req := httptest.NewRequest(
		http.MethodGet, "/students/1", nil)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d",
			resp.StatusCode)
	}

	// 检查响应头
	if got := resp.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("unexpected Content-Type: %s", got)
	}

	// 解析 JSON 响应体
	var got Student
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response failed: %v", err)
	}

	want := Student{
		ID:   1,
		Name: "Eva",
		Age:  23,
	}

	if got != want {
		t.Errorf("expected %+v, got %+v", want, got)
	}
}

// 测试不存在的学生
func TestGetStudentNotFound(t *testing.T) {
	handler := newStudentHandler()

	req := httptest.NewRequest(
		http.MethodGet, "/students/999", nil)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d",
			resp.StatusCode)
	}
}
func TestGetStudentBadRequest(t *testing.T) {
	handler := newStudentHandler()

	req := httptest.NewRequest(
		http.MethodGet, "/students/abc", nil)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf(
			"expected status 400, got %d",
			resp.StatusCode,
		)
	}
}
