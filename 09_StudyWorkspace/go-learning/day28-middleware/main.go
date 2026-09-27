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
	log.Println(
		"handler",
	)

	fmt.Fprintln(
		w,
		"hello from handler",
	)
}

func middlewareA(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			log.Println(
				"A: before",
			)

			next.ServeHTTP(
				w,
				r,
			)

			log.Println(
				"A: after",
			)
		},
	)
}

func middlewareB(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			log.Println(
				"B: before",
			)

			next.ServeHTTP(
				w,
				r,
			)

			log.Println(
				"B: after",
			)
		},
	)
}

func main() {
	mux := http.NewServeMux()

	businessHandler :=
		http.HandlerFunc(
			helloHandler,
		)

	wrappedHandler :=
		middlewareA(
			middlewareB(
				businessHandler,
			),
		)

	mux.Handle(
		"/hello",
		wrappedHandler,
	)

	log.Println(
		"server started at http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			mux,
		),
	)
}
