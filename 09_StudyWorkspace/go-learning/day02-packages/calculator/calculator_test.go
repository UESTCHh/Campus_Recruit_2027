package calculator

import "testing"

func TestAdd(t *testing.T) {
	got := Add(10, 5)
	want := 15

	if got != want {
		t.Fatalf("Add() = %d, want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(10, 5)
	want := 5

	if got != want {
		t.Fatalf("Subtract() = %d, want %d", got, want)
	}
}
