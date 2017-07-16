package math32

import (
	"math"
	"testing"
)

func TestSinh(t *testing.T) { testOneFloatToOneFloat(t, "Sinh", Sinh, math.Sinh) }

func TestCosh(t *testing.T) { testOneFloatToOneFloat(t, "Cosh", Cosh, math.Cosh) }
