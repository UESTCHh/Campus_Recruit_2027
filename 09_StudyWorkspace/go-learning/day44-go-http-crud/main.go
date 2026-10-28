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

// PUT /students/{id}
// 根据路径中的ID更新学生姓名和年龄
func updateStudentHandler(repo *StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 解析路径中的学生ID
		studentID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || studentID <= 0 {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid student id"},
			)
			return
		}

		// 2. 定义PUT请求体
		// ID不从请求体读取，而是由URL路径确定
		var req struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		// 3. 解码JSON
		decoder := json.NewDecoder(r.Body)
		// 它使解码器在遇到结构体不认识的 JSON 字段时返回错误
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&req); err != nil {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid JSON body"},
			)
			return
		}

		// 4. 校验数据
		if strings.TrimSpace(req.Name) == "" || req.Age <= 0 {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid student data"},
			)
			return
		}

		// 5. 根据路径ID构造待更新的学生记录
		student := Student{
			ID:   studentID,
			Name: req.Name,
			Age:  req.Age,
		}

		// 6. 调用Repository
		affected, err := repo.Update(student)

		if err != nil {
			log.Printf("Update failed: %v", err)

			writeJSON(
				w,
				http.StatusInternalServerError,
				ErrorResponse{Error: "internal server error"},
			)
			return
		}

		// 7. 判断学生是否存在
		if affected == 0 {
			writeJSON(
				w,
				http.StatusNotFound,
				ErrorResponse{Error: "student not found"},
			)
			return
		}

		// 8. 返回更新后的学生信息
		writeJSON(w, http.StatusOK, student)
	}
}

// DELETE /students/{id}
// 根据学生ID删除指定记录
func deleteStudentHandler(repo *StudentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 从URL路径获取ID
		idText := r.PathValue("id")

		// 2. 转换并校验ID
		studentID, err := strconv.Atoi(idText)
		if err != nil || studentID <= 0 {
			writeJSON(
				w,
				http.StatusBadRequest,
				ErrorResponse{Error: "invalid student id"},
			)
			return
		}

		// 3. 调用Repository删除学生
		affected, err := repo.Delete(studentID)
		if err != nil {
			log.Printf("Delete failed: %v", err)

			writeJSON(
				w,
				http.StatusInternalServerError,
				ErrorResponse{Error: "internal server error"},
			)
			return
		}

		// 4. 目标学生不存在
		if affected == 0 {
			writeJSON(
				w,
				http.StatusNotFound,
				ErrorResponse{Error: "student not found"},
			)
			return
		}

		// 5. 删除成功，不返回JSON响应体
		w.WriteHeader(http.StatusNoContent)
	}
}

// 路由注册函数
func registerStudentRoutes(
	mux *http.ServeMux,
	repo *StudentRepository,
) {
	// TODO：
	// 将 main() 中原来的五条 mux.HandleFunc()
	// 原样移动到这里
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

	mux.HandleFunc(
		"PUT /students/{id}",
		updateStudentHandler(repo),
	)

	mux.HandleFunc(
		"DELETE /students/{id}",
		deleteStudentHandler(repo),
	)
}

func main() {
	// // 复用之前的文件型SQLite数据库
	// dbPath := filepath.Join(
	// 	"..",
	// 	"day38-sql-persistence",
	// 	"students.db",
	// )

	// 1. 尝试从环境变量 STUDENT_DB_PATH 读取数据库路径
	dbPath := os.Getenv("STUDENT_DB_PATH")
	// 2. 如果没有配置，则使用原来的默认数据库路径
	if dbPath == "" {
		dbPath = filepath.Join(
			"..",
			"day38-sql-persistence",
			"students.db",
		)
	}
	// 3. 保留 os.Stat() 文件存在性检查
	// 防止意外创建新的空数据库
	// 因为 SQLite 在指定文件不存在时可能创建新的空数据库
	// 而我们当前不希望因为数据库路径写错，就意外创建一个没有 students 表的文件。
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

	// 数据库连接正常后，确保 students 表存在
	if err := initStudentTable(db); err != nil {
		log.Fatalf("initialize students table failed: %v", err)
	}

	// 复用现有Repository
	repo := NewStudentRepository(db)

	mux := http.NewServeMux()

	registerStudentRoutes(mux, repo)

	log.Println("Server listening on http://127.0.0.1:8081")

	if err := http.ListenAndServe("127.0.0.1:8081", mux); err != nil {
		log.Fatal(err)
	}
}
