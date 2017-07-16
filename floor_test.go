package math32

import (
	"math"
	"testing"
)

func TestFloor(t *testing.T) { testOneFloatToOneFloat(t, "Floor", Floor, math.Floor) }

func TestCeil(t *testing.T) { testOneFloatToOneFloat(t, "Ceil", Ceil, math.Ceil) }

func TestTrunc(t *testing.T) { testOneFloatToOneFloat(t, "Trunc", Trunc, math.Trunc) }
