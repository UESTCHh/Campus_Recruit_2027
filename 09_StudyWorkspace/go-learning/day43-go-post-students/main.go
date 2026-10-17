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
	"strings"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// JSON错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// 统一写入JSON响应
func writeJSON(w http.ResponseWriter, status int, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("JSON encoding failed: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if _, err := w.Write(append(body, '\n')); err != nil {
		log.Printf("write response failed: %v", err)
	}
}

// GET /students/{id}
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

// POST /students
func createStudentHandler(repo *StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 读取并解码请求JSON
		var student Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid JSON body"},
			)
			return
		}

		// 2. 校验学生数据
		if student.ID <= 0 ||
			strings.TrimSpace(student.Name) == "" ||
			student.Age <= 0 {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid student data"},
			)
			return
		}

		// 3. 调用Repository写入SQLite
		err = repo.Create(student)
		if err != nil {
			// 判断是否属于SQLite的主键或唯一约束冲突
			var sqliteErr sqlite3.Error

			if errors.As(err, &sqliteErr) &&
				(sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey ||
					sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique) {
				writeJSON(
					w,
					http.StatusConflict,
					ErrorResponse{Error: "student id already exists"},
				)
				return
			}

			// 其他数据库错误
			log.Printf("Create failed: %v", err)

			writeJSON(
				w,
				http.StatusInternalServerError,
				ErrorResponse{Error: "internal server error"},
			)
			return
		}

		// 4. 新增成功：设置资源地址
		w.Header().Set(
			"Location",
			"/students/"+strconv.Itoa(student.ID),
		)

		// 5. 返回201及新建学生
		writeJSON(w, http.StatusCreated, student)
	}
}

func main() {
	// 复用之前的文件型SQLite数据库
	dbPath := filepath.Join(
		"..",
		"day38-sql-persistence",
		"students.db",
	)

	// 防止意外创建新的空数据库
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

	// 复用现有Repository
	repo := NewStudentRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /students",
		studentsHandler(repo),
	)

	mux.HandleFunc(
		"GET /students/{id}",
		studentByIDHandler(repo),
	)

	mux.HandleFunc(
		"POST /students",
		createStudentHandler(repo),
	)

	log.Println("Server listening on http://127.0.0.1:8081")

	if err := http.ListenAndServe("127.0.0.1:8081", mux); err != nil {
		log.Fatal(err)
	}
}
