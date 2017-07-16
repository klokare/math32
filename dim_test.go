package math32

import (
	"math"
	"testing"
)

func TestDim(t *testing.T) { testTwoFloatToOneFloat(t, "Dim", Dim, math.Dim) }

func TestMin(t *testing.T) { testTwoFloatToOneFloat(t, "Min", Min, math.Min) }

func TestMax(t *testing.T) { testTwoFloatToOneFloat(t, "Max", Max, math.Max) }
