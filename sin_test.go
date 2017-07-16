package math32

import (
	"math"
	"testing"
)

func TestSin(t *testing.T) { testOneFloatToOneFloat(t, "Sin", Sin, math.Sin) }

func TestCos(t *testing.T) { testOneFloatToOneFloat(t, "Cos", Cos, math.Cos) }
