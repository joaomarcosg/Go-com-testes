package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	iterations := Iteration("a", 5)
	expected := "aaaaa"

	if iterations != expected {
		t.Errorf("expected '%s', but got '%s'", expected, iterations)
	}
}

func BenchMarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a", 5)
	}
}

func ExampleIteration() {
	iteration := Iteration("a", 5)
	fmt.Println(iteration)
	// Output: aaaaa
}
