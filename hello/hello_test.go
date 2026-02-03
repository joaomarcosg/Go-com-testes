package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("João", "")
		expected := "Olá, João"
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say 'Olá, Mundo' when a empty string is receveid", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá, Mundo"
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		result := Hello("Elodie", "spanish")
		expected := "Hola, Elodie"
		verifyCorrectMessage(t, result, expected)
	})

}
