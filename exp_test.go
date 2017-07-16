package math32

import (
	"math"
	"testing"
)

func TestExp(t *testing.T) { testOneFloatToOneFloat(t, "Exp", Exp, math.Exp) }

func TestExp2(t *testing.T) { testOneFloatToOneFloat(t, "Exp2", Exp2, math.Exp2) }
