package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		f()
	}
}
