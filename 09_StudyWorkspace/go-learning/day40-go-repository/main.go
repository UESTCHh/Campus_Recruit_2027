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

func main() {
	// 1. 定位Day38的数据库文件
	dbPath := filepath.Join(
		"..",
		"day38-sql-persistence",
		"students.db",
	)

	// 避免数据库文件缺失时创建空数据库
	if _, err := os.Stat(dbPath); err != nil {
		log.Fatalf("database file not found: %v", err)
	}

	// 2. 创建数据库连接池管理对象
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 3. 检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 4. 创建Repository
	repo := NewStudentRepository(db)

	// 第一部分：根据ID查询一条学生记录
	fmt.Println("GetByID:")

	studentIDs := []int{1001, 1002, 9999}

	for _, studentID := range studentIDs {
		student, err := repo.GetByID(studentID)

		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Student %d not found\n", studentID)
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"%d | %s | %d\n",
			student.ID,
			student.Name,
			student.Age,
		)
	}

	// 第二部分：查询全部学生记录
	fmt.Println("\nList:")

	students, err := repo.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		fmt.Printf(
			"%d | %s | %d\n",
			student.ID,
			student.Name,
			student.Age,
		)
	}
}
