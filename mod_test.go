package math32

import (
	"math"
	"testing"
)

func TestMod(t *testing.T) { testTwoFloatToOneFloat(t, "Mod", Mod, math.Mod) }
