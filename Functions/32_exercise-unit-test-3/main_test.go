package main

import (
	"log"
	"testing"
)

func TestLandscape(t *testing.T) {
	got := landscape("Bali")
	want := "My idea of landscape is Bali"
	if got != want {
		log.Fatalf("error - want: %s and got: %s", want, got)
	}
}
