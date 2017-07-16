package math32

import (
	"math"
	"testing"
)

func TestExpm1(t *testing.T) { testOneFloatToOneFloat(t, "Expm1", Expm1, math.Expm1) }
