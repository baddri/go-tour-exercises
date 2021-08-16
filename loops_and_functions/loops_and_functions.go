package loops_and_functions

import (
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for x != toFixed(z*z, 5) {
		z -= (z*z - x) / (2 * z)
	}
	return toFixed(z, 5)
}

// see https://stackoverflow.com/a/29786394/12953170
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
