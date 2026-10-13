package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// 注册SQLite驱动
	_ "github.com/mattn/go-sqlite3"
)

// Student表示一名学生
type Student struct {
	ID   int
	Name string
	Age  int
}

// findStudentByID根据学生ID查询单条记录
// db：共享的数据库管理对象
// studentID：需要查询的学生ID
// 返回Student和error
func findStudentByID(db *sql.DB, studentID int) (Student, error) {
	var student Student

	// 使用参数化查询，避免直接拼接SQL
	err := db.QueryRow(
		"SELECT id, name, age FROM students WHERE id = ?",
		studentID,
	).Scan(
		&student.ID,
		&student.Name,
		&student.Age,
	)

	// 将查询错误返回给调用者
	if err != nil {
		return Student{}, err
	}

	return student, nil
}

func main() {
	// 定位Day38创建的SQLite数据库文件
	dbPath := filepath.Join(
		"..",
		"day38-sql-persistence",
		"students.db",
	)

	// 确认数据库文件已经存在
	if _, err := os.Stat(dbPath); err != nil {
		log.Fatalf("database file not found: %v", err)
	}

	// 创建数据库连接池管理对象
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 验证数据库连接可用
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 模拟查询三个学生ID
	studentIDs := []int{1001, 1002, 9999}

	for _, studentID := range studentIDs {
		// 调用封装好的查询函数
		student, err := findStudentByID(db, studentID)

		// 学生不存在：单独处理
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Student %d not found\n", studentID)
			continue
		}

		// 其他数据库错误
		if err != nil {
			log.Fatal(err)
		}

		// 查询成功
		fmt.Printf(
			"%d | %s | %d\n",
			student.ID,
			student.Name,
			student.Age,
		)
	}
}
