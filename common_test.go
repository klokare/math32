package math32

import (
	"math"
	"testing"
)

// Test values
var values = []float32{
	// Not a number
	NaN(),
	// ±Infinity
	Inf(-1),
	Inf(1),
	// Really big
	MaxFloat32,
	-MaxFloat32,
	// Really small but non-zero
	SmallestNonzeroFloat32,
	-SmallestNonzeroFloat32,
	// Numbers bigger than 1
	5, 2, 1,
	-5, -2, -1,
	// Numbers smaller than 1
	0.9, 0.5, 0.2, 0.1,
	-0.9, -0.5, -0.2, -0.1,
	// Zero
	0, -0,
}

func testOneFloatToOneFloat(t *testing.T, name string, comp32 func(float32) float32, comp64 func(float64) float64) {

	for _, x := range values {
		actual := comp32(x)
		expected := comp64(float64(x))

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
			t.Errorf("%s(%f) expected NaN, actual %f", name, x, actual)

		case (!math.IsNaN(expected) && IsNaN(actual)):
			t.Errorf("%s(%f) expected %f, actual NaN", name, x, expected)

		case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
			t.Errorf("%s(%f) expected -Inf, actual %f", name, x, actual)

		case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
			t.Errorf("%s(%f) expected %f, actual -Inf", name, x, expected)

		case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
			t.Errorf("%s(%f) expected Inf, actual %f", name, x, actual)

		case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
			t.Errorf("%s(%f) expected %f, actual Inf", name, x, expected)

		default:
			if actual != float32(expected) {
				t.Errorf("%s(%f) expected %f, actual %f", name, x, expected, actual)
			}
		}
	}
}

func testTwoFloatToOneFloat(t *testing.T, name string, comp32 func(float32, float32) float32, comp64 func(float64, float64) float64) {

	for _, x := range values {
		for _, y := range values {
			actual := comp32(x, y)
			expected := comp64(float64(x), float64(y))

			switch {
			case (math.IsNaN(expected) && IsNaN(actual)),
				(math.IsInf(expected, -1) && IsInf(actual, -1)),
				(math.IsInf(expected, 1) && IsInf(actual, 1)),
				(expected > float64(MaxFloat32) && IsInf(actual, 1)),
				(expected < float64(-MaxFloat32) && IsInf(actual, -1)):
				// All is good

			case (math.IsNaN(expected) && !IsNaN(actual)):
				t.Errorf("%s(%f, %f) expected NaN, actual %f", name, x, y, actual)

			case (!math.IsNaN(expected) && IsNaN(actual)):
				t.Errorf("%s(%f, %f) expected %f, actual NaN", name, x, y, expected)

			case (math.IsInf(expected, -1) && !IsInf(actual, -1)):
				t.Errorf("%s(%f, %f) expected -Inf, actual %f", name, x, y, actual)

			case (!math.IsInf(expected, -1) && IsInf(actual, -1)):
				t.Errorf("%s(%f, %f) expected %f, actual -Inf", name, x, y, expected)

			case (math.IsInf(expected, 1) && !IsInf(actual, 1)):
				t.Errorf("%s(%f, %f) expected Inf, actual %f", name, x, y, actual)

			case (!math.IsInf(expected, 1) && IsInf(actual, 1)):
				t.Errorf("%s(%f, %f) expected %f, actual Inf", name, x, y, expected)

			default:
				if actual != float32(expected) {
					t.Errorf("%s(%f, %f) expected %f, actual %f", name, x, y, expected, actual)
				}
			}
		}
	}
}
