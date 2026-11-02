// Derived from hqx-java.
// Copyright 2003 Maxim Stepin, 2010 Cameron Zemek,
// 2011 Tamme Schichler, and 2012 A. Eduardo García.
// Licensed under the GNU Lesser General Public License, version 3 or later.
// See LGPL-3.0.txt.

// Package rgbYuv converts packed RGB colors to the YUV values used by HQX.
package rgbYuv

// GetYuv converts a packed RGB color to the YUV format used by HQX.
// The alpha channel is ignored.
func GetYuv(rgb int) int {
	c := uint32(rgb)
	r, g, b := float64((c>>16)&255), float64((c>>8)&255), float64(c&255)
	y := int(0.299*r + 0.587*g + 0.114*b)
	u := int(-0.169*r-0.331*g+0.500*b) + 128
	v := int(0.500*r-0.419*g-0.081*b) + 128
	return y<<16 | u<<8 | v
}

// HqxInit is retained for compatibility. Conversion no longer needs a lookup table.
func HqxInit() {}
