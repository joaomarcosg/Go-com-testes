package iteration

import "testing"

func TestIteration(t *testing.T) {
	iterations := Iteration("a")
	expected := "aaaaa"

	if iterations != expected {
		t.Errorf("expected '%s', but got '%s'", expected, iterations)
	}
}

func BenchMarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a")
	}
}
