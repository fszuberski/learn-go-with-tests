package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Filip")

	got := buffer.String()
	want := "Hello, Filip"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
