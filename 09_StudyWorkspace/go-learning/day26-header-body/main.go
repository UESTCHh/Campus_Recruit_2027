package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func headerHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Content-Type 用于描述请求Body的数据格式。
	contentType :=
		r.Header.Get("Content-Type")

	// User-Agent通常用于描述发起请求的客户端信息。
	userAgent :=
		r.Header.Get("User-Agent")

	// 教学用自定义Header。
	studentName :=
		r.Header.Get("X-Student-Name")

	fmt.Fprintln(
		w,
		"Content-Type:",
		contentType,
	)

	fmt.Fprintln(
		w,
		"User-Agent:",
		userAgent,
	)

	fmt.Fprintln(
		w,
		"X-Student-Name:",
		studentName,
	)

	fmt.Fprintln(
		w,
		"Host:",
		r.Host,
	)
}

func bodyHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// r.Body是请求主体的数据流。
	bodyBytes, err :=
		io.ReadAll(r.Body)

	if err != nil {
		http.Error(
			w,
			"failed to read body",
			http.StatusBadRequest,
		)
		return
	}

	// io.ReadAll得到[]byte。
	// 当前实验为了观察内容，将其转换成string。
	fmt.Fprintln(
		w,
		"body:",
		string(bodyBytes),
	)
}

func createUserHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 用struct承接客户端提交的JSON数据。
	var req CreateUserRequest

	// Decoder直接从Request Body数据流读取JSON，
	// 并把结果写入req。
	err :=
		json.NewDecoder(r.Body).
			Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid json body",
			http.StatusBadRequest,
		)
		return
	}

	fmt.Fprintln(
		w,
		"name:",
		req.Name,
	)

	fmt.Fprintln(
		w,
		"age:",
		req.Age,
	)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/headers",
		headerHandler,
	)

	mux.HandleFunc(
		"/body",
		bodyHandler,
	)

	mux.HandleFunc(
		"/users",
		createUserHandler,
	)

	log.Println(
		"starting server at http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			mux,
		),
	)
}
