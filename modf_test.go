package math32

import (
	"math"
	"testing"
)

func TestModf(t *testing.T) {
	name := "Modf"

	for _, x := range values {
		actualInt, actualFrac := Modf(x)
		expectedInt, expectedFrac := math.Modf(float64(x))

		// Compare Int values
		switch {
		case (math.IsNaN(expectedInt) && IsNaN(actualInt)),
			(math.IsInf(expectedInt, -1) && IsInf(actualInt, -1)),
			(math.IsInf(expectedInt, 1) && IsInf(actualInt, 1)),
			(expectedInt > float64(MaxFloat32) && IsInf(actualInt, 1)),
			(expectedInt < float64(-MaxFloat32) && IsInf(actualInt, -1)):
			// All is good

		case IsInf(actualInt, 1) && (float32(expectedInt) == actualInt):
			// All is good

		case (math.IsNaN(expectedInt) && !IsNaN(actualInt)):
			t.Errorf("%s(%f) expectedInt NaN, actualInt %f", name, x, actualInt)

		case (!math.IsNaN(expectedInt) && IsNaN(actualInt)):
			t.Errorf("%s(%f) expectedInt %f, actualInt NaN", name, x, expectedInt)

		case (math.IsInf(expectedInt, -1) && !IsInf(actualInt, -1)):
			t.Errorf("%s(%f) expectedInt -Inf, actualInt %f", name, x, actualInt)

		case (!math.IsInf(expectedInt, -1) && IsInf(actualInt, -1)):
			t.Errorf("%s(%f) expectedInt %f, actualInt -Inf", name, x, expectedInt)

		case (math.IsInf(expectedInt, 1) && !IsInf(actualInt, 1)):
			t.Errorf("%s(%f) expectedInt Inf, actualInt %f", name, x, actualInt)

		case (!math.IsInf(expectedInt, 1) && IsInf(actualInt, 1)):
			t.Errorf("%s(%f) expectedInt %f, actualInt Inf", name, x, expectedInt)

		default:
			if actualInt != float32(expectedInt) {
				t.Errorf("%s(%f) expectedInt %f, actualInt %f", name, x, expectedInt, actualInt)
			}
		}

		// Compare Frac values
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
	}

}
