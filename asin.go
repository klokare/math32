// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Asin returns the arcsine, in radians, of x.
//
// Special cases are:
//	Asin(±0) = ±0
//	Asin(x) = NaN if x < -1 or x > 1
func Asin(x float32) float32 {
	return float32(math.Asin(float64(x)))
}

// Acos returns the arccosine, in radians, of x.
//
// Special case is:
//	Acos(x) = NaN if x < -1 or x > 1
func Acos(x float32) float32 {
	return float32(math.Acos(float64(x)))
}
