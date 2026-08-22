package main

import "testing"

// *testing.T 是 Go 测试框架交给这个测试函数的“测试控制对象”。
func TestGreet(t *testing.T) {
	got := greet("hui")
	want := "Hello, hui!"

	if got != want {
		t.Fatalf("greet() = %q, want %q", got, want)
	}
}

func TestClassifyScore(t *testing.T) {
	got := classifyScore(85)
	want := "passed"

	if got != want {
		t.Fatalf("classifyScore() = %q, want %q", got, want)
	}
}

func TestSumScores(t *testing.T) {
	scores := []int{80, 90, 70}

	got := sumScores(scores)
	want := 240

	if got != want {
		t.Fatalf("sumScores() = %d, want %d", got, want)
	}
}
