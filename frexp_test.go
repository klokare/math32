package math32

import (
	"math"
	"testing"
)

func TestFrexp(t *testing.T) {
	name := "Frexp"
	for _, x := range values {
		actualFrac, actualExp := Frexp(x)
		expectedFrac, expectedExp := math.Frexp(float64(x))

		switch {
		case (math.IsNaN(expectedFrac) && IsNaN(actualFrac)),
			(math.IsInf(expectedFrac, -1) && IsInf(actualFrac, -1)),
			(math.IsInf(expectedFrac, 1) && IsInf(actualFrac, 1)),
			(expectedFrac > float64(MaxFloat32) && IsInf(actualFrac, 1)),
			(expectedFrac < float64(-MaxFloat32) && IsInf(actualFrac, -1)):
			// All is good

		case IsInf(actualFrac, 1) && (float32(expectedFrac) == actualFrac):
			// All is good

		case (math.IsNaN(expectedFrac) && !IsNaN(actualFrac)):
			t.Errorf("%s(%f) expectedFrac NaN, actualFrac %f", name, x, actualFrac)

		case (!math.IsNaN(expectedFrac) && IsNaN(actualFrac)):
			t.Errorf("%s(%f) expectedFrac %f, actualFrac NaN", name, x, expectedFrac)

		case (math.IsInf(expectedFrac, -1) && !IsInf(actualFrac, -1)):
			t.Errorf("%s(%f) expectedFrac -Inf, actualFrac %f", name, x, actualFrac)

		case (!math.IsInf(expectedFrac, -1) && IsInf(actualFrac, -1)):
			t.Errorf("%s(%f) expectedFrac %f, actualFrac -Inf", name, x, expectedFrac)

		case (math.IsInf(expectedFrac, 1) && !IsInf(actualFrac, 1)):
			t.Errorf("%s(%f) expectedFrac Inf, actualFrac %f", name, x, actualFrac)

		case (!math.IsInf(expectedFrac, 1) && IsInf(actualFrac, 1)):
			t.Errorf("%s(%f) expectedFrac %f, actualFrac Inf", name, x, expectedFrac)

		default:
			if actualFrac != float32(expectedFrac) {
				t.Errorf("%s(%f) expectedFrac %f, actualFrac %f", name, x, expectedFrac, actualFrac)
			}
		}

		if actualExp != expectedExp {
			t.Errorf("%s(%f) expectedFrac %d, actualExp %d", name, x, expectedFrac, actualExp)
		}
	}
}
