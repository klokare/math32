// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Sinh returns the hyperbolic sine of x.
//
// Special cases are:
//	Sinh(±0) = ±0
//	Sinh(±Inf) = ±Inf
//	Sinh(NaN) = NaN
func Sinh(x float32) float32 {
	return float32(math.Sinh(float64(x)))
}

// Cosh returns the hyperbolic cosine of x.
//
// Special cases are:
//	Cosh(±0) = 1
//	Cosh(±Inf) = +Inf
//	Cosh(NaN) = NaN
func Cosh(x float32) float32 {
	return float32(math.Cosh(float64(x)))
}
