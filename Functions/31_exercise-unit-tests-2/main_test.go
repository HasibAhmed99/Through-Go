package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(7, 10)
	want := 17
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 10)
	want := -3
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(15, 5, add)
	want := 20
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}
