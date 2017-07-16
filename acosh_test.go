package math32

import (
	"math"
	"testing"
)

func TestAcosh(t *testing.T) { testOneFloatToOneFloat(t, "Acosh", Acosh, math.Acosh) }
