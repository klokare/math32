package math32

import (
	"math"
	"testing"
)

func TestJ1(t *testing.T) { testOneFloatToOneFloat(t, "J1", J1, math.J1) }

func TestY1(t *testing.T) { testOneFloatToOneFloat(t, "Y1", Y1, math.Y1) }
