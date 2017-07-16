package math32

import (
	"math"
	"testing"
)

func TestJ0(t *testing.T) { testOneFloatToOneFloat(t, "J0", J0, math.J0) }

func TestY0(t *testing.T) { testOneFloatToOneFloat(t, "Y0", Y0, math.Y0) }
