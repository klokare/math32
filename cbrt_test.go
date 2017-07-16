package math32

import (
	"math"
	"testing"
)

func TestCbrt(t *testing.T) { testOneFloatToOneFloat(t, "Cbrt", Cbrt, math.Cbrt) }
