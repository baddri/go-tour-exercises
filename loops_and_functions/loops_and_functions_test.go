package loops_and_functions

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	if Sqrt(16) != math.Sqrt(16) {
		t.Fatalf("Sqrt(16) = %v, want match for %v", Sqrt(16), math.Sqrt(16))
	}
}
