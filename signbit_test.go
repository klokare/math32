package math32

import (
	"math"
	"testing"
)

func TestSignbit(t *testing.T) {
	name := "Signbit"
	for _, x := range values {
		actual := Signbit(x)
		expected := math.Signbit(float64(x))
		if actual != expected {
			t.Errorf("%s(%f) expected %v, actual %v", name, x, expected, actual)
		}
	}
}
