package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// API错误响应的统一结构。
//
// Code：
// 给客户端程序判断错误类型。
//
// Message：
// 给人阅读的错误描述。
type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeError(
	w http.ResponseWriter,
	status int,
	code string,
	message string,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(
		errorResponse{
			Code:    code,
			Message: message,
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
			http.StatusBadRequest,
			"INVALID_NAME",
			"name is required",
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
