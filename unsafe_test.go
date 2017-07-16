package math32

import (
	"math"
	"testing"
)

func TestUnsafe32(t *testing.T) {
	for _, x := range values {
		actualB := Float32bits(x)
		expectedB := math.Float32bits(x)
		if actualB != expectedB {
			t.Errorf("Float32bits(%f) expected %v, actual %v", x, expectedB, actualB)
		}
		actualF := Float32frombits(expectedB)
		expectedF := math.Float32frombits(expectedB)
		if !(IsNaN(actualF) && IsNaN(expectedF)) && actualF != expectedF {
			t.Errorf("Float32frombits(%v) expected %f, actual %f", expectedB, expectedF, actualF)
		}
	}
}

func TestUnsafe64(t *testing.T) {
	for _, x32 := range values {
		x := float64(x32)
		actualB := Float64bits(x)
		expectedB := math.Float64bits(x)
		if actualB != expectedB {
			t.Errorf("Float64bits(%f) expected %v, actual %v", x, expectedB, actualB)
		}
		actualF := Float64frombits(expectedB)
		expectedF := math.Float64frombits(expectedB)
		if !(math.IsNaN(actualF) && math.IsNaN(expectedF)) && actualF != expectedF {
			t.Errorf("Float64frombits(%v) expected %f, actual %f", expectedB, expectedF, actualF)
		}
	}
}
