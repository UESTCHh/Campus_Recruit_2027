package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(
			w,
			"name is required",
			http.StatusBadRequest,
		)

		// 故意不写return
		// return
	}

	fmt.Fprintln(
		w,
		"hello",
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
