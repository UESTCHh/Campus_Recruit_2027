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

	// // 3. 定义建表 SQL
	// createSQL := `
	//     CREATE TABLE IF NOT EXISTS students (
	//         id INTEGER PRIMARY KEY,
	//         name TEXT NOT NULL,
	//         age INTEGER NOT NULL
	//     );
	// `

	// // TODO 1：使用 db.Exec(createSQL) 第一次建表
	// // 必须检查 error
	// _, err = db.Exec(createSQL)
	// if err != nil {
	// 	t.Fatalf("create table failed: %v", err)
	// }
	err = initStudentTable(db)
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
	err = initStudentTable(db)
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

// func TestInitStudentTableError(t *testing.T) {
// 	// TODO 1：使用 sql.Open() 打开一个临时 SQLite 数据库
// 	// 1. 创建临时数据库路径
// 	dbPath := filepath.Join(t.TempDir(), "day49_test.db")

//		// 2. 打开 SQLite 数据库
//		db, err := sql.Open("sqlite3", dbPath)
//		// TODO 2：检查打开数据库时的 error
//		if err != nil {
//			t.Fatal(err)
//		}
//		t.Cleanup(func() {
//			if err := db.Close(); err != nil {
//				t.Errorf("close database: %v", err)
//			}
//		})
//		// TODO 3：调用 db.Close() 主动关闭数据库，并检查错误
//		err = db.Close()
//		if err != nil {
//			t.Errorf("close database: %v", err)
//		}
//		// TODO 4：调用 initStudentTable(db)
//		err = initStudentTable(db)
//		// TODO 5：验证返回的 error 不为 nil
//		if err == nil {
//			t.Fatalf("expected error but got nil")
//		}
//	}

// 这段整理与上面我写的测试逻辑一致，只删除了不必要的重复关闭。
func TestInitStudentTableError(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "closed_test.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatal(err)
	}

	// 主动关闭数据库，制造初始化失败条件
	if err := db.Close(); err != nil {
		t.Fatalf("close database: %v", err)
	}

	// 关闭后执行建表，应返回非 nil 的 error
	err = initStudentTable(db)
	if err == nil {
		t.Fatal("expected error for closed database, got nil")
	}
}
