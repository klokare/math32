package math32

import (
	"math"
	"testing"
)

func TestAtan2(t *testing.T) { testTwoFloatToOneFloat(t, "Atan2", Atan2, math.Atan2) }
