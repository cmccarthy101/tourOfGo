package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	test := 999999.0
	x, _ := Sqrt(test)
	actual := math.Sqrt(test)
	diff := math.Abs(x - actual)
	tolerance := 0.000000001
	if diff > tolerance {
		t.Errorf("expected a diff of %v, but got: %v", tolerance, diff)
	}
}
