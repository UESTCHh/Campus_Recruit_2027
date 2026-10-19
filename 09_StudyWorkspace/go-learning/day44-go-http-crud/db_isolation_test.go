package main

import (
	"database/sql"
	"errors"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// newIsolatedStudentRepo 为每次调用创建独立的 SQLite 测试数据库。
// 仅供测试使用，不会访问 Day38 的 students.db。
func newIsolatedStudentRepo(t *testing.T) *StudentRepository {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "students_test.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("open test database: %v", err)
	}

	// t.TempDir() 的清理先注册；这里的 Cleanup 后注册，先关闭数据库再删目录。
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("close test database: %v", err)
		}
	})

	db.SetMaxOpenConns(1)
	if _, err := db.Exec(`
		CREATE TABLE students (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		)
	`); err != nil {
		t.Fatalf("create test table: %v", err)
	}

	return NewStudentRepository(db)
}

func TestStudentRepositoryIsolation(t *testing.T) {
	repoA := newIsolatedStudentRepo(t)
	repoB := newIsolatedStudentRepo(t)

	want := Student{ID: 9001, Name: "Eva", Age: 23}
	if err := repoA.Create(want); err != nil {
		t.Fatalf("create student in database A: %v", err)
	}

	got, err := repoA.GetByID(9001)
	if err != nil {
		t.Fatalf("get student from database A: %v", err)
	}
	if got != want {
		t.Errorf("database A: want %+v, got %+v", want, got)
	}

	_, err = repoB.GetByID(9001)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("database B: want sql.ErrNoRows, got %v", err)
	}
}
