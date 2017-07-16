package math32

import (
	"math"
	"testing"
)

func TestNextAfter32(t *testing.T) {
	for _, x := range values {
		for _, y := range values {
			actual := Nextafter32(x, y)
			expected := math.Nextafter32(x, y)
			if !(IsNaN(actual) && IsNaN(expected)) && actual != expected {
				t.Errorf("Nextafter32(%f,%f) expected %f, actual %f", x, y, expected, actual)
			}
		}
	}
}

func TestNextAfter(t *testing.T) {
	for _, x32 := range values {
		x := float64(x32)
		for _, y32 := range values {
			y := float64(y32)
			actual := Nextafter(x, y)
			expected := math.Nextafter(x, y)
			if !(math.IsNaN(actual) && math.IsNaN(expected)) && actual != expected {
				t.Errorf("Nextafter(%f,%f) expected %f, actual %f", x, y, expected, actual)
			}
		}
	}
}
