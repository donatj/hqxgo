// Derived from hqx-java.
// Copyright 2003 Maxim Stepin, 2010 Cameron Zemek,
// 2011 Tamme Schichler, and 2012 A. Eduardo García.
// Licensed under the GNU Lesser General Public License, version 3 or later.
// See LGPL-3.0.txt.

// Package interpolation blends packed ARGB colors.
package interpolation

// mix blends each channel independently. The weights must add up to a power of two.
func mix(c1, c2, c3 int, w1, w2, w3 uint32, shift uint) int {
	a, b, c := uint32(c1), uint32(c2), uint32(c3)
	var result uint32
	for _, n := range [...]uint{0, 8, 16, 24} {
		value := (((a>>n)&255)*w1 + ((b>>n)&255)*w2 + ((c>>n)&255)*w3) >> shift
		result |= value << n
	}
	return int(result)
}

func Mix3To1(c1, c2 int) int        { return mix(c1, c2, 0, 3, 1, 0, 2) }
func Mix2To1To1(c1, c2, c3 int) int { return mix(c1, c2, c3, 2, 1, 1, 2) }
func Mix7To1(c1, c2 int) int        { return mix(c1, c2, 0, 7, 1, 0, 3) }
func Mix2To7To7(c1, c2, c3 int) int { return mix(c1, c2, c3, 2, 7, 7, 4) }
func MixEven(c1, c2 int) int        { return mix(c1, c2, 0, 1, 1, 0, 1) }

// Mix4To2To1 uses weights 5:2:1, matching the original implementation.
func Mix4To2To1(c1, c2, c3 int) int  { return mix(c1, c2, c3, 5, 2, 1, 3) }
func Mix6To1To1(c1, c2, c3 int) int  { return mix(c1, c2, c3, 6, 1, 1, 3) }
func Mix5To3(c1, c2 int) int         { return mix(c1, c2, 0, 5, 3, 0, 3) }
func Mix2To3To3(c1, c2, c3 int) int  { return mix(c1, c2, c3, 2, 3, 3, 3) }
func Mix14To1To1(c1, c2, c3 int) int { return mix(c1, c2, c3, 14, 1, 1, 4) }
