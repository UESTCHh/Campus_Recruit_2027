package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// 统一JSON错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// 统一返回JSON内容
func writeJSON(w http.ResponseWriter, status int, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("JSON encoding failed: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if _, err := w.Write(append(body, '\n')); err != nil {
		log.Printf("write response failed: %v", err)
	}
}

// GET /students/{id}
// 根据ID查询单个学生
func studentByIDHandler(repo *StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idText := r.PathValue("id")

		studentID, err := strconv.Atoi(idText)
		if err != nil || studentID <= 0 {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid student id"},
			)
			return
		}

		student, err := repo.GetByID(studentID)

		if errors.Is(err, sql.ErrNoRows) {
			writeJSON(
				w,
				http.StatusNotFound,
				ErrorResponse{Error: "student not found"},
			)
			return
		}

		if err != nil {
			log.Printf("GetByID failed: %v", err)

			writeJSON(
				w,
				http.StatusInternalServerError,
				ErrorResponse{Error: "internal server error"},
			)
			return
		}

		writeJSON(w, http.StatusOK, student)
	}
}

// GET /students
// 查询全部学生
func studentsHandler(repo *StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := repo.List()
		if err != nil {
			log.Printf("List failed: %v", err)

			writeJSON(
				w,
				http.StatusInternalServerError,
				ErrorResponse{Error: "internal server error"},
			)
			return
		}

		writeJSON(w, http.StatusOK, students)
	}
}

func main() {
	dbPath := filepath.Join(
		"..",
		"day38-sql-persistence",
		"students.db",
	)

	if _, err := os.Stat(dbPath); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := NewStudentRepository(db)

	mux := http.NewServeMux()

	// 查询全部学生
	mux.HandleFunc(
		"GET /students",
		studentsHandler(repo),
	)

	// 根据ID查询学生
	mux.HandleFunc(
		"GET /students/{id}",
		studentByIDHandler(repo),
	)

	log.Println("Server listening on http://127.0.0.1:8080")

	if err := http.ListenAndServe("127.0.0.1:8080", mux); err != nil {
		log.Fatal(err)
	}
}
