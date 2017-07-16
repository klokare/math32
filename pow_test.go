package math32

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) { testTwoFloatToOneFloat(t, "Pow", Pow, math.Pow) }
