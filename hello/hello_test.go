package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Leo")
	want := "Hello, Leo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
