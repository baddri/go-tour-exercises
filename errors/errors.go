package errors

import (
	"fmt"

	"github.com/baddri/go-tour-exercises/loops_and_functions"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func SqrtWithError(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return loops_and_functions.Sqrt(x), nil
}
