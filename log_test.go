package math32

import (
	"math"
	"testing"
)

func TestLog(t *testing.T) { testOneFloatToOneFloat(t, "Log", Log, math.Log) }
