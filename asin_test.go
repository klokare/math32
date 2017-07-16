package math32

import (
	"math"
	"testing"
)

func TestAsin(t *testing.T) { testOneFloatToOneFloat(t, "Asin", Asin, math.Asin) }

func TestAcos(t *testing.T) { testOneFloatToOneFloat(t, "Acos", Acos, math.Acos) }
