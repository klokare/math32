package math32

import (
	"math"
	"testing"
)

func TestAsinh(t *testing.T) { testOneFloatToOneFloat(t, "Asinh", Asinh, math.Asinh) }
