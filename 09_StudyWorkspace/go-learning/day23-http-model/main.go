package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler负责处理 /hello 请求。
func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(
		w,
		"hello",
	)
}

// nameHandler负责处理 /name 请求。
func nameHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(
		w,
		"my name is hui",
	)
}

// studyHandler负责处理 /study 请求。
func studyHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println(
		"studyHandler is running",
	)

	fmt.Fprintln(
		w,
		"I am studying Go backend",
	)
}
func main() {
	// 当Client访问 /hello，
	// Server就把Request交给helloHandler。
	http.HandleFunc(
		"/hello",
		helloHandler,
	)

	// 当Client访问 /name，
	// Server就把Request交给nameHandler。
	http.HandleFunc(
		"/name",
		nameHandler,
	)

	// 当Client访问 /study，
	// Server就把Request交给studyHandler。
	http.HandleFunc(
		"/study",
		studyHandler,
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
