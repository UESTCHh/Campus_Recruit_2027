package main

import (
	"database/sql"
	"errors"
	"strings"
	"testing"
)

func TestGetByIDErrorWrapping(t *testing.T) {
	// TODO 1：创建独立 Repository
	repo := newIsolatedStudentRepo(t)
	// TODO 2：查询一个不存在的学生，例如 9999
	_, err := repo.GetByID(9999)
	// TODO 3：断言 err 不为 nil
	if err == nil {
		t.Fatalf("expected error but got nil")
	}
	// TODO 4：使用 errors.Is()
	// 验证是否仍然能够识别 sql.ErrNoRows
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("test failed")
	}
	// TODO 5：使用 strings.Contains()
	// 验证错误信息是否包含 "GetByID"
	if !strings.Contains(err.Error(), "GetByID") {
		t.Fatalf("expected error to contain GetByID, got: %v", err)
	}
}
