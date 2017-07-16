package math32

import (
	"math"
	"testing"
)

func TestPow10(t *testing.T) {
	name := "Pow10"

	for _, e := range []int{-325, -324, -100, -5, -2, -1, -0, 0, 1, 2, 5, 100, 309, 310} {
		actual := Pow10(e)
		expected := math.Pow10(e)

		switch {
		case (math.IsNaN(expected) && IsNaN(actual)),
			(math.IsInf(expected, -1) && IsInf(actual, -1)),
			(math.IsInf(expected, 1) && IsInf(actual, 1)),
			(expected > float64(MaxFloat32) && IsInf(actual, 1)),
			(expected < float64(-MaxFloat32) && IsInf(actual, -1)):
			// All is good

		case IsInf(actual, 1) && (float32(expected) == actual):
			// All is good

		case (math.IsNaN(expected) && !IsNaN(actual)):
			t.Errorf("%s(%d) expected NaN, actual %f", name, e, actual)

		case (!math.IsNaN(expected) && IsNaN(actual)):
			t.Errorf("%s(%d) expected %f, actual NaN", name, e, expected)

		case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
			t.Errorf("%s(%d) expected -Inf, actual %f", name, e, actual)

		case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
			t.Errorf("%s(%d) expected %f, actual -Inf", name, e, expected)

		case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
			t.Errorf("%s(%d) expected Inf, actual %f", name, e, actual)

		case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
			t.Errorf("%s(%d) expected %f, actual Inf", name, e, expected)

		default:
			if actual != float32(expected) {
				t.Errorf("%s(%d) expected %f, actual %f", name, e, expected, actual)
			}
		}
	}
}
