package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

// Context Key 使用专用类型，避免 Key 冲突。
type requestIDKey struct{}

// 使用原子计数器，为每个请求分配教学用 Request ID。
var requestCounter atomic.Uint64

// APIError：服务器内部完整的错误描述。
type APIError struct {
	Status  int
	Code    string
	Message string
}

// errorResponse：实际返回客户端的 JSON Body。
type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 预定义错误值。
var ErrInvalidName = APIError{
	Status:  http.StatusBadRequest,
	Code:    "INVALID_NAME",
	Message: "name is required",
}

// Middleware：生成 Request ID，并传递给下游 Handler。
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		id := fmt.Sprintf(
			"req-%d",
			requestCounter.Add(1),
		)

		// 保存 Request ID 到 Context。
		ctx := context.WithValue(
			r.Context(),
			requestIDKey{},
			id,
		)

		// 将 Request ID 返回给客户端。
		w.Header().Set("X-Request-ID", id)

		// 继续执行下游 Handler。
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// writeError：统一记录错误日志，并返回 JSON 错误响应。
//
// 与 Stage 1 相比，新增 r *http.Request，
// 用于读取当前请求 Context 中的 Request ID。
func writeError(
	w http.ResponseWriter,
	r *http.Request,
	apiErr APIError,
) {
	// 从请求 Context 中读取 Middleware 保存的 ID。
	requestID := r.Context().Value(requestIDKey{})

	// 服务端错误日志：
	// 同时记录 Request ID 和 APIError 信息。
	log.Printf(
		"request_id=%v status=%d code=%s message=%q",
		requestID,
		apiErr.Status,
		apiErr.Code,
		apiErr.Message,
	)

	// 下面仍然沿用 Day34 的统一 JSON Error 处理。
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(apiErr.Status)

	json.NewEncoder(w).Encode(
		errorResponse{
			Code:    apiErr.Code,
			Message: apiErr.Message,
		},
	)
}

// Handler：只判断请求是否合法，并选择对应错误。
func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get("name")

	if name == "" {
		writeError(w, r, ErrInvalidName)
		return
	}

	fmt.Fprintf(w, "hello %s\n", name)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	// 请求必须先经过 Request ID Middleware。
	handler := requestIDMiddleware(mux)

	log.Println("server started at http://localhost:8080")

	log.Fatal(
		http.ListenAndServe(":8080", handler),
	)
}
