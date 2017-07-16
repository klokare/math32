package math32

import (
	"math"
	"testing"
)

func TestAtan(t *testing.T) { testOneFloatToOneFloat(t, "Atan", Atan, math.Atan) }
