package main

import (
	"errors"
	"testing"
)

func TestGetUserNameSuccess(t *testing.T) {
	name, err := GetUserName(1)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if name != "UESTCHh" {
		t.Fatalf("expected UESTCHh, got %s", name)
	}
}

func TestGetUserNameNotFound(t *testing.T) {
	_, err := GetUserName(100)

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}
