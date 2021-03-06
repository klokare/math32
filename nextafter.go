// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

// Nextafter32 returns the next representable float32 value after x towards y.
//
// Special cases are:
//	Nextafter32(x, x)   = x
//	Nextafter32(NaN, y) = NaN
//	Nextafter32(x, NaN) = NaN
func Nextafter32(x, y float32) (r float32) {
	return math.Nextafter32(x, y)
}

// Nextafter returns the next representable float64 value after x towards y.
//
// Special cases are:
//	Nextafter(x, x)   = x
//	Nextafter(NaN, y) = NaN
//	Nextafter(x, NaN) = NaN
func Nextafter(x, y float64) (r float64) {
	return math.Nextafter(x, y)
}
