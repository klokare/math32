package math32

import (
	"math"
	"testing"
)

func TestLog1p(t *testing.T) { testOneFloatToOneFloat(t, "Log1p", Log1p, math.Log1p) }
