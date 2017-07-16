package math32

import (
	"math"
	"testing"
)

func TestInf(t *testing.T) {
	actual := Inf(1)
	if !IsInf(actual, 1) {
		t.Errorf("Inf(1) expected +Inf, actual %f", actual)
	}
	actual = Inf(-1)
	if !IsInf(actual, -1) {
		t.Errorf("Inf(-1) expected -Inf, actual %f", actual)
	}
}

func TestNaN(t *testing.T) {
	actual := NaN()
	if !IsNaN(actual) {
		t.Errorf("NaN() expected NaN, actual %f", actual)
	}
}

func TestIsNaN(t *testing.T) {
	name := "IsNaN"
	for _, x := range values {
		actual := IsNaN(x)
		expected := math.IsNaN(float64(x))
		if actual != expected {
			t.Errorf("%s(%f) expected %v, actual %v", name, x, expected, actual)
		}
	}
}

func TestIsInf(t *testing.T) {
	name := "IsInf"
	for _, x := range values {
		for _, sign := range []int{-100, -1, 0, 1, 100} {
			actual := IsInf(x, sign)
			expected := math.IsInf(float64(x), sign)
			if actual != expected {
				t.Errorf("%s(%f, %d) expected %v, actual %v", name, x, sign, expected, actual)
			}
		}
	}
}
