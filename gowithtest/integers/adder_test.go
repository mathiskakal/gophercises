package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// adding the output like so as a comment makes it an example test, without it it won't run
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
