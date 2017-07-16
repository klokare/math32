package math32

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	testOneFloatToOneFloat(t, "Abs", Abs, math.Abs)
}
