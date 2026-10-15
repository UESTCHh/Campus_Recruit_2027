package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 使用内存SQLite数据库，不修改原有students.db
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 避免不同连接指向不同的内存数据库
	// 是为了保证本实验中的数据库操作复用同一个底层 SQLite 连接
	// 避免多个连接各自访问不同的内存数据库。
	db.SetMaxOpenConns(1)

	// 初始化学生表
	_, err = db.Exec(`
		CREATE TABLE students (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := NewStudentRepository(db)

	// 1. Create：新增学生
	err = repo.Create(Student{
		ID:   1004,
		Name: "David",
		Age:  23,
	})
	if err != nil {
		log.Fatal(err)
	}

	student, err := repo.GetByID(1004)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("After Create: %d | %s | %d\n",
		student.ID, student.Name, student.Age)

	// 2. Update：修改年龄
	affected, err := repo.Update(Student{
		ID:   1004,
		Name: "David",
		Age:  24,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Update affected: %d\n", affected)

	student, err = repo.GetByID(1004)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("After Update: %d | %s | %d\n",
		student.ID, student.Name, student.Age)

	// 3. Update：尝试修改不存在的学生
	affected, err = repo.Update(Student{
		ID:   9999,
		Name: "Nobody",
		Age:  20,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Update missing affected: %d\n", affected)

	// 4. Delete：删除存在的学生
	affected, err = repo.Delete(1004)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Delete affected: %d\n", affected)

	// 5. Delete：再次删除同一学生
	affected, err = repo.Delete(1004)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Delete again affected: %d\n", affected)

	// 6. 验证删除结果
	_, err = repo.GetByID(1004)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Student 1004 not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal("Student 1004 should have been deleted")
	}
}
