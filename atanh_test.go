package math32

import (
	"math"
	"testing"
)

func TestAtanh(t *testing.T) { testOneFloatToOneFloat(t, "Atanh", Atanh, math.Atanh) }
