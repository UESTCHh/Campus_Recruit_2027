package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type contextKey string

const userIDKey contextKey = "userID"

// A：产生userID并写入Context。
func userMiddleware(
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

			userID := 1001

			ctx := context.WithValue(
				r.Context(),
				userIDKey,
				userID,
			)

			r = r.WithContext(ctx)

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

// B：不创建userID，只读取上游传下来的userID。
func auditMiddleware(
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

			userID, ok :=
				r.Context().
					Value(userIDKey).(int)

			if !ok {
				http.Error(
					w,
					"user id not found",
					http.StatusInternalServerError,
				)
				return
			}

			log.Println(
				"B: userID =",
				userID,
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

func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Println(
		"handler: running",
	)

	userID, ok :=
		r.Context().
			Value(userIDKey).(int)

	if !ok {
		http.Error(
			w,
			"user id not found",
			http.StatusInternalServerError,
		)
		return
	}

	log.Println(
		"handler: userID =",
		userID,
	)

	fmt.Fprintf(
		w,
		"hello user %d\n",
		userID,
	)
}

func main() {
	mux := http.NewServeMux()

	businessHandler :=
		http.HandlerFunc(
			helloHandler,
		)

	wrappedHandler :=
		userMiddleware(
			auditMiddleware(
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
