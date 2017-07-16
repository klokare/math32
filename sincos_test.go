package math32

import (
	"math"
	"testing"
)

func TestSincos(t *testing.T) {
	name := "Sincos"

	for _, x := range values {
		actualSin, actualCos := Sincos(x)
		expectedSin, expectedCos := math.Sincos(float64(x))

		// Compare Sin values
		switch {
		case (math.IsNaN(expectedSin) && IsNaN(actualSin)),
			(math.IsInf(expectedSin, -1) && IsInf(actualSin, -1)),
			(math.IsInf(expectedSin, 1) && IsInf(actualSin, 1)),
			(expectedSin > float64(MaxFloat32) && IsInf(actualSin, 1)),
			(expectedSin < float64(-MaxFloat32) && IsInf(actualSin, -1)):
			// All is good

		case IsInf(actualSin, 1) && (float32(expectedSin) == actualSin):
			// All is good

		case (math.IsNaN(expectedSin) && !IsNaN(actualSin)):
			t.Errorf("%s(%f) expectedSin NaN, actualSin %f", name, x, actualSin)

		case (!math.IsNaN(expectedSin) && IsNaN(actualSin)):
			t.Errorf("%s(%f) expectedSin %f, actualSin NaN", name, x, expectedSin)

		case (math.IsInf(expectedSin, -1) && !IsInf(actualSin, -1)):
			t.Errorf("%s(%f) expectedSin -Inf, actualSin %f", name, x, actualSin)

		case (!math.IsInf(expectedSin, -1) && IsInf(actualSin, -1)):
			t.Errorf("%s(%f) expectedSin %f, actualSin -Inf", name, x, expectedSin)

		case (math.IsInf(expectedSin, 1) && !IsInf(actualSin, 1)):
			t.Errorf("%s(%f) expectedSin Inf, actualSin %f", name, x, actualSin)

		case (!math.IsInf(expectedSin, 1) && IsInf(actualSin, 1)):
			t.Errorf("%s(%f) expectedSin %f, actualSin Inf", name, x, expectedSin)

		default:
			if actualSin != float32(expectedSin) {
				t.Errorf("%s(%f) expectedSin %f, actualSin %f", name, x, expectedSin, actualSin)
			}
		}

		// Compare Cos values
		switch {
		case (math.IsNaN(expectedCos) && IsNaN(actualCos)),
			(math.IsInf(expectedCos, -1) && IsInf(actualCos, -1)),
			(math.IsInf(expectedCos, 1) && IsInf(actualCos, 1)),
			(expectedCos > float64(MaxFloat32) && IsInf(actualCos, 1)),
			(expectedCos < float64(-MaxFloat32) && IsInf(actualCos, -1)):
			// All is good

		case IsInf(actualCos, 1) && (float32(expectedCos) == actualCos):
			// All is good

		case (math.IsNaN(expectedCos) && !IsNaN(actualCos)):
			t.Errorf("%s(%f) expectedCos NaN, actualCos %f", name, x, actualCos)

		case (!math.IsNaN(expectedCos) && IsNaN(actualCos)):
			t.Errorf("%s(%f) expectedCos %f, actualCos NaN", name, x, expectedCos)

		case (math.IsInf(expectedCos, -1) && !IsInf(actualCos, -1)):
			t.Errorf("%s(%f) expectedCos -Inf, actualCos %f", name, x, actualCos)

		case (!math.IsInf(expectedCos, -1) && IsInf(actualCos, -1)):
			t.Errorf("%s(%f) expectedCos %f, actualCos -Inf", name, x, expectedCos)

		case (math.IsInf(expectedCos, 1) && !IsInf(actualCos, 1)):
			t.Errorf("%s(%f) expectedCos Inf, actualCos %f", name, x, actualCos)

		case (!math.IsInf(expectedCos, 1) && IsInf(actualCos, 1)):
			t.Errorf("%s(%f) expectedCos %f, actualCos Inf", name, x, expectedCos)

		default:
			if actualCos != float32(expectedCos) {
				t.Errorf("%s(%f) expectedCos %f, actualCos %f", name, x, expectedCos, actualCos)
			}
		}
	}

}
