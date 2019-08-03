package iteration

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got: '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("c", 8))
	// Output: cccccccc
}
