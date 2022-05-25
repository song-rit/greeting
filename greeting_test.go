package geekgreeting

import (
	"testing"
)

// func TestAbs(t *testing.T) {
// 	got := Abs(-1)
// 	// if got != 1 {
// 	// 	t.Errorf("Abs(-1) = %d; want 1", got)
// 	// }
// }

func TestHello(t *testing.T) {
	hello := "Hello"
	if out := Hello(); out != hello {
		t.Errorf("Hello != %q, want %q", out, hello)
	}
}
