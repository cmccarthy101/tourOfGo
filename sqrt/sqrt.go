package sqrt

import (
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	diff := 99.0
	i := 0

	for diff > 0.0000001 {
		i++
		diff = z
		z = z - (z*z-x)/(2*z)

		diff = math.Abs(diff - z)
	}

	return z, i
}