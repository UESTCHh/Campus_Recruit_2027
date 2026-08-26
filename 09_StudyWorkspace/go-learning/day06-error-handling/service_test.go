package main

import (
	"errors"
	"testing"
)

func TestGetUserDisplayNameSuccess(t *testing.T) {
	name, err := GetUserDisplayName(1)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if name != "User:UESTCHh" {
		t.Fatalf(
			"expected User:UESTCHh, got %s",
			name,
		)
	}
}

func TestGetUserDisplayNameNotFound(t *testing.T) {
	_, err := GetUserDisplayName(100)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected wrapped ErrUserNotFound, got %v",
			err,
		)
	}
}
