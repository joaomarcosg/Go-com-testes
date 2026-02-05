package iteration

import "testing"

func TestIteration(t *testing.T) {
	iterations := Iteration("a")
	expected := "aaaaa"

	if iterations != expected {
		t.Errorf("expected '%s', but got '%s'", expected, iterations)
	}
}
