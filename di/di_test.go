package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gopher")

	got := buffer.String()
	expected := "Hello, Gopher"

	assertEquals(t, got, expected)
}

func assertEquals(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}
