package math32

import (
	"math"
	"testing"
)

func TestTanh(t *testing.T) { testOneFloatToOneFloat(t, "Tanh", Tanh, math.Tanh) }
