package math32

import (
	"math"
	"testing"
)

func TestLogb(t *testing.T) { testOneFloatToOneFloat(t, "Logb", Logb, math.Logb) }

func TestIlogb(t *testing.T) {
	for _, x := range values {
		actual := Ilogb(x)
		expected := math.Ilogb(float64(x))
		if actual != expected {
			t.Errorf("Ilogb(%f) expected %d, actual %d", x, expected, actual)
		}
	}
}
