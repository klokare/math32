package math32

import (
	"math"
	"testing"
)

func TestJn(t *testing.T) {
	name := "Jn"

	for _, n := range []int{-5, -2, -1, 0, 1, 2, 5} {
		for _, x := range values {
			actual := Jn(n, x)
			expected := math.Jn(n, float64(x))

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
				t.Errorf("%s(%d,%f) expected NaN, actual %f", name, n, x, actual)

			case (!math.IsNaN(expected) && IsNaN(actual)):
				t.Errorf("%s(%d,%f) expected %f, actual NaN", name, n, x, expected)

			case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
				t.Errorf("%s(%d,%f) expected -Inf, actual %f", name, n, x, actual)

			case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
				t.Errorf("%s(%d,%f) expected %f, actual -Inf", name, n, x, expected)

			case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
				t.Errorf("%s(%d,%f) expected Inf, actual %f", name, n, x, actual)

			case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
				t.Errorf("%s(%d,%f) expected %f, actual Inf", name, n, x, expected)

			default:
				if actual != float32(expected) {
					t.Errorf("%s(%d, %f) expected %f, actual %f", name, n, x, expected, actual)
				}
			}
		}
	}
}

func TestYn(t *testing.T) {
	name := "Yn"

	for _, n := range []int{-5, -2, -1, 0, 1, 2, 5} {
		for _, x := range values {
			actual := Yn(n, x)
			expected := math.Yn(n, float64(x))

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
				t.Errorf("%s(%d,%f) expected NaN, actual %f", name, n, x, actual)

			case (!math.IsNaN(expected) && IsNaN(actual)):
				t.Errorf("%s(%d,%f) expected %f, actual NaN", name, n, x, expected)

			case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
				t.Errorf("%s(%d,%f) expected -Inf, actual %f", name, n, x, actual)

			case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
				t.Errorf("%s(%d,%f) expected %f, actual -Inf", name, n, x, expected)

			case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
				t.Errorf("%s(%d,%f) expected Inf, actual %f", name, n, x, actual)

			case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
				t.Errorf("%s(%d,%f) expected %f, actual Inf", name, n, x, expected)

			default:
				if actual != float32(expected) {
					t.Errorf("%s(%d, %f) expected %f, actual %f", name, n, x, expected, actual)
				}
			}
		}
	}
}
