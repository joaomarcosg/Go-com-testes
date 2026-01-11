package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("João")
	expected := "Hello, João"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
