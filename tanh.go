// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Tanh returns the hyperbolic tangent of x.
//
// Special cases are:
//	Tanh(±0) = ±0
//	Tanh(±Inf) = ±1
//	Tanh(NaN) = NaN
func Tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}
