package math32

import (
	"math"
	"testing"
)

func TestCopysign(t *testing.T) { testTwoFloatToOneFloat(t, "Copysign", Copysign, math.Copysign) }
