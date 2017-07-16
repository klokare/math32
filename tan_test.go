package math32

import (
	"math"
	"testing"
)

func TestTan(t *testing.T) { testOneFloatToOneFloat(t, "Tan", Tan, math.Tan) }
