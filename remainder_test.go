package math32

import (
	"math"
	"testing"
)

func TestRemainder(t *testing.T) { testTwoFloatToOneFloat(t, "Remainder", Remainder, math.Remainder) }
