// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Jn returns the order-n Bessel function of the first kind.
//
// Special cases are:
//	Jn(n, ±Inf) = 0
//	Jn(n, NaN) = NaN
func Jn(n int, x float32) float32 {
	return float32(math.Jn(n, float64(x)))
}

// Yn returns the order-n Bessel function of the second kind.
//
// Special cases are:
//	Yn(n, +Inf) = 0
//	Yn(n > 0, 0) = -Inf
//	Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
//	Y1(n, x < 0) = NaN
//	Y1(n, NaN) = NaN
func Yn(n int, x float32) float32 {
	return float32(math.Yn(n, float64(x)))
}
