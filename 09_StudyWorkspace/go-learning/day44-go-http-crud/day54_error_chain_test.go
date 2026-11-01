package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestDay54ErrorWrapping(t *testing.T) {
	// TODO 1：定义原始错误 sql.ErrNoRows
	err1 := sql.ErrNoRows
	// TODO 2：使用 fmt.Errorf + %w 包装错误
	err2 := fmt.Errorf("query failed: %w", err1)
	// TODO 3：使用 fmt.Errorf + %v 创建另一个错误
	err3 := fmt.Errorf("query failed: %v", err1)
	// TODO 4：通过 errors.Is 检查两个错误
	// %w：应该为 true
	// %v：应该为 false
	if !errors.Is(err2, sql.ErrNoRows) {
		t.Fatalf("expected true, but got false")
	}
	if errors.Is(err3, sql.ErrNoRows) {
		t.Fatalf("expected false, but got true")
	}

	// TODO 5：验证 errors.Unwrap 的结果
	err4 := errors.Unwrap(err2)
	if err4 != sql.ErrNoRows {
		t.Fatalf("expected sql.ErrNoRows, but got %v", err4)
	}
}

func TestDay54ErrorsAs(t *testing.T) {
	pathErr := &os.PathError{
		Op:   "open",
		Path: "students.db",
		Err:  os.ErrNotExist,
	}

	err := fmt.Errorf("initialize database failed: %w", pathErr)

	var target *os.PathError

	if !errors.As(err, &target) {
		t.Fatalf("expected *os.PathError in error chain, got %v", err)
	}

	if target.Op != "open" || target.Path != "students.db" {
		t.Fatalf(
			"unexpected PathError: Op=%q, Path=%q",
			target.Op,
			target.Path,
		)
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatal("wrong error variable")
	}
}
