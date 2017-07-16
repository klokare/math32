package math32

import (
	"math"
	"testing"
)

func TestHypot(t *testing.T) { testTwoFloatToOneFloat(t, "Hypot", Hypot, math.Hypot) }
