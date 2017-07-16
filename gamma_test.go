package math32

import (
	"math"
	"testing"
)

func TestGamma(t *testing.T) { testOneFloatToOneFloat(t, "Gamma", Gamma, math.Gamma) }
