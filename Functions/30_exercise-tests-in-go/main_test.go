package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	// we can also do like this:
	/* total := Add(5, 5)
	expected := 10
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, expected)
	} */
}
