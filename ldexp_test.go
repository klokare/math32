package math32

import (
	"math"
	"testing"
)

func TestLdexp(t *testing.T) {
	name := "Ldexp"

	for _, exp := range []int{-5, -2, -1, 0, 1, 2, 5} {
		for _, frac := range values {
			actual := Ldexp(frac, exp)
			expected := math.Ldexp(float64(frac), exp)

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
				t.Errorf("%s(%f,%f) expected NaN, actual %f", name, frac, exp, actual)

			case (!math.IsNaN(expected) && IsNaN(actual)):
				t.Errorf("%s(%f,%f) expected %f, actual NaN", name, frac, exp, expected)

			case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
				t.Errorf("%s(%f,%f) expected -Inf, actual %f", name, frac, exp, actual)

			case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
				t.Errorf("%s(%f,%f) expected %f, actual -Inf", name, frac, exp, expected)

			case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
				t.Errorf("%s(%f,%f) expected Inf, actual %f", name, frac, exp, actual)

			case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
				t.Errorf("%s(%f,%f) expected %f, actual Inf", name, frac, exp, expected)

			default:
				if actual != float32(expected) {
					t.Errorf("%s(%f,%f) expected %f, actual %f", name, frac, exp, expected, actual)
				}
			}
		}
	}
}
