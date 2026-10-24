package main

import (
	"database/sql"
	"path/filepath"
	"testing"
)

func TestDay49SQLiteInitialization(t *testing.T) {
	// 1. 创建临时数据库路径
	dbPath := filepath.Join(t.TempDir(), "day49_test.db")

	// 2. 打开 SQLite 数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("close database: %v", err)
		}
	})

	// 3. 定义建表 SQL
	createSQL := `
        CREATE TABLE IF NOT EXISTS students (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            age INTEGER NOT NULL
        );
    `

	// TODO 1：使用 db.Exec(createSQL) 第一次建表
	// 必须检查 error
	_, err = db.Exec(createSQL)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	// TODO 2：插入学生 (1001, "Alice", 20)
	// 可以使用 db.Exec()，记得检查 error
	_, err = db.Exec("INSERT INTO students(id, name, age) VALUES(?, ?, ?)", 1001, "Alice", 20)
	if err != nil {
		t.Fatalf("INSERT student failed: %v", err)
	}
	// TODO 3：第二次执行 db.Exec(createSQL)
	// 仍然必须检查 error
	_, err = db.Exec(createSQL)
	if err != nil {
		t.Fatalf("create table failed: %v", err)
	}
	// TODO 4：查询 ID=1001 的学生
	// 提示：db.QueryRow() + row.Scan()
	var got Student
	err = db.QueryRow("SELECT * FROM students where id = ?", 1001).Scan(&got.ID, &got.Name, &got.Age)
	if err != nil {
		t.Fatalf("query student failed: %v", err)
	}

	// TODO 5：断言查询结果
	// Name 应为 Alice，Age 应为 20
	want := Student{1001, "Alice", 20}
	if got != want {
		t.Fatalf("expected %v but got %v", want, got)
	}
}
