package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// APIError 表示服务器内部的一次完整 API 错误。
type APIError struct {
	Status  int
	Code    string
	Message string
}

// errorResponse 是真正写入 Response Body 的 JSON 结构。
type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 常见 API 错误集中定义。
// Handler 不需要每次重复填写 Status / Code / Message。
var ErrInvalidName = APIError{
	Status:  http.StatusBadRequest,
	Code:    "INVALID_NAME",
	Message: "name is required",
}

func writeError(
	w http.ResponseWriter,
	apiErr APIError,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		apiErr.Status,
	)

	json.NewEncoder(w).Encode(
		errorResponse{
			Code:    apiErr.Code,
			Message: apiErr.Message,
		},
	)
}

func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	name :=
		r.URL.Query().Get("name")

	if name == "" {
		writeError(
			w,
			ErrInvalidName,
		)

		return
	}

	fmt.Fprintf(
		w,
		"hello %s\n",
		name,
	)
}

func main() {
	http.HandleFunc(
		"/hello",
		helloHandler,
	)

	log.Println(
		"server started at http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			nil,
		),
	)
}
