package math32

import (
	"math"
	"testing"
)

func TestLgamma(t *testing.T) {
	name := "Lgamma"
	for _, x := range values {
		actualLgamma, actualSign := Lgamma(x)
		expectedLgamma, expectedSign := math.Lgamma(float64(x))

		switch {
		case (math.IsNaN(expectedLgamma) && IsNaN(actualLgamma)),
			(math.IsInf(expectedLgamma, -1) && IsInf(actualLgamma, -1)),
			(math.IsInf(expectedLgamma, 1) && IsInf(actualLgamma, 1)),
			(expectedLgamma > float64(MaxFloat32) && IsInf(actualLgamma, 1)),
			(expectedLgamma < float64(-MaxFloat32) && IsInf(actualLgamma, -1)):
			// All is good

		case IsInf(actualLgamma, 1) && (float32(expectedLgamma) == actualLgamma):
			// All is good

		case (math.IsNaN(expectedLgamma) && !IsNaN(actualLgamma)):
			t.Errorf("%s(%f) expectedLgamma NaN, actualLgamma %f", name, x, actualLgamma)

		case (!math.IsNaN(expectedLgamma) && IsNaN(actualLgamma)):
			t.Errorf("%s(%f) expectedLgamma %f, actualLgamma NaN", name, x, expectedLgamma)

		case (math.IsInf(expectedLgamma, -1) && !IsInf(actualLgamma, -1)):
			t.Errorf("%s(%f) expectedLgamma -Inf, actualLgamma %f", name, x, actualLgamma)

		case (!math.IsInf(expectedLgamma, -1) && IsInf(actualLgamma, -1)):
			t.Errorf("%s(%f) expectedLgamma %f, actualLgamma -Inf", name, x, expectedLgamma)

		case (math.IsInf(expectedLgamma, 1) && !IsInf(actualLgamma, 1)):
			t.Errorf("%s(%f) expectedLgamma Inf, actualLgamma %f", name, x, actualLgamma)

		case (!math.IsInf(expectedLgamma, 1) && IsInf(actualLgamma, 1)):
			t.Errorf("%s(%f) expectedLgamma %f, actualLgamma Inf", name, x, expectedLgamma)

		default:
			if actualLgamma != float32(expectedLgamma) {
				t.Errorf("%s(%f) expectedLgamma %f, actualLgamma %f", name, x, expectedLgamma, actualLgamma)
			}
		}

		if actualSign != expectedSign {
			t.Errorf("%s(%f) expectedLgamma %d, actualSign %d", name, x, expectedLgamma, actualSign)
		}
	}
}
