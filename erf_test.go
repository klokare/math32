package math32

import (
	"math"
	"testing"
)

func TestErf(t *testing.T) { testOneFloatToOneFloat(t, "Erf", Erf, math.Erf) }

func TestErfc(t *testing.T) { testOneFloatToOneFloat(t, "Erfc", Erfc, math.Erfc) }
