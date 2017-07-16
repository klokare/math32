package math32

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) { testOneFloatToOneFloat(t, "Sqrt", Sqrt, math.Sqrt) }
