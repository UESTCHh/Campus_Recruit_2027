package main

import (
	"database/sql"
)

// initStudentTable 负责确保 students 表存在。
// 成功返回 nil，失败返回 error。
func initStudentTable(db *sql.DB) error {
	// TODO 1：复制 Day49 的 CREATE TABLE IF NOT EXISTS SQL
	createSQL := `
		CREATE TABLE IF NOT EXISTS students(
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		);
	`

	// TODO 2：使用 db.Exec() 执行 SQL
	_, err := db.Exec(createSQL)

	// TODO 3：检查 err，失败时返回错误
	if err != nil {
		return err
	}
	// TODO 4：成功时返回 nil
	return nil
}
