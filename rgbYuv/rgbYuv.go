package rgbYuv

var rgbMask int = 16777215 //0x00FFFFFF;
var RGBtoYUV [16777216]int //0x1000000

/**
 * Returns the 24bit YUV equivalent of the provided 24bit RGB color.<b>Any alpha component is dropped</b>
 *
 * @param rgb a 24bit rgb color
 * @return the corresponding 24bit YUV color
 */
func GetYuv(rgb int) int {
	return RGBtoYUV[rgb&rgbMask]
}

/**
 * Calculates the lookup table. <b>MUST</b> be called (only once) before doing anything else
 */
func HqxInit() {
	/* Initalize RGB to YUV lookup table */
	var r, g, b, y, u, v int
	for c := 16777216 - 1; c >= 0; c-- {
		r = (c & 16711680) >> 16
		g = (c & 65280) >> 8
		b = c & 255
		y = int(+0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
		u = int(-0.169*float32(r)-0.331*float32(g)+0.500*float32(b)) + 128
		v = int(+0.500*float32(r)-0.419*float32(g)-0.081*float32(b)) + 128
		RGBtoYUV[c] = (y << 16) | (u << 8) | v
	}
}
