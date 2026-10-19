package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func newStudentHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students/{id}",
		func(w http.ResponseWriter, r *http.Request) {

			id, err := strconv.Atoi(r.PathValue("id"))
			if err != nil || id <= 0 {
				http.Error(w, "invalid student id",
					http.StatusBadRequest)
				return
			}

			if id != 1 {
				http.Error(w, "student not found",
					http.StatusNotFound)
				return
			}

			student := Student{
				ID:   1,
				Name: "Eva",
				Age:  23,
			}

			w.Header().Set(
				"Content-Type", "application/json")

			json.NewEncoder(w).Encode(student)
		})

	return mux
}
