package errors

import "testing"

func TestSqrtWithError(t *testing.T) {
	t.Run(
		"Positive number shouldn't return error",
		func(t *testing.T) {
			v, err := SqrtWithError(8)
			if err != nil {
				t.Fatalf("Sqrt(8) = %v, %v want match for %v, <nil>", v, err, v)
			}
		},
	)
	t.Run(
		"Negative number should return error",
		func(t *testing.T) {
			v, err := SqrtWithError(-8)
			if err == nil {
				t.Fatalf("Sqrt(-8) = %v, %v want match for %v, Error('cannot Sqrt negative number'))", v, err, v)
			}
		},
	)
}
