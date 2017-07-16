package math32

import (
	"math"
	"testing"
)

func TestLog10(t *testing.T) { testOneFloatToOneFloat(t, "Log10", Log10, math.Log10) }

func TestLog2(t *testing.T) { testOneFloatToOneFloat(t, "Log2", Log2, math.Log2) }
