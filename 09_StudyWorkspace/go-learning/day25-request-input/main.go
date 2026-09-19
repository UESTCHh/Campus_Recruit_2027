package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func inspectHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(
		w,
		"method:",
		r.Method,
	)

	fmt.Fprintln(
		w,
		"path:",
		r.URL.Path,
	)

	fmt.Fprintln(
		w,
		"raw query:",
		r.URL.RawQuery,
	)

	fmt.Fprintln(
		w,
		"host:",
		r.Host,
	)

	fmt.Fprintln(
		w,
		"protocol:",
		r.Proto,
	)
}
func searchHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query()

	keyword :=
		query.Get("keyword")

	pageText :=
		query.Get("page")

	// Required Check：
	// page在这个Demo中必须提供。
	if pageText == "" {
		http.Error(
			w,
			"page is required",
			http.StatusBadRequest,
		)

		return
	}

	// Parse：
	// Query Parameter默认是string，
	// 业务需要int，所以进行转换。
	page, err :=
		strconv.Atoi(pageText)

	if err != nil {
		http.Error(
			w,
			"page must be an integer",
			http.StatusBadRequest,
		)

		return
	}

	// Validation：
	// 类型正确不代表业务值一定合理。
	if page < 1 {
		http.Error(
			w,
			"page must be at least 1",
			http.StatusBadRequest,
		)

		return
	}

	fmt.Fprintln(
		w,
		"keyword:",
		keyword,
	)

	fmt.Fprintln(
		w,
		"page:",
		page,
	)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/inspect",
		inspectHandler,
	)

	mux.HandleFunc(
		"/search",
		searchHandler,
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
