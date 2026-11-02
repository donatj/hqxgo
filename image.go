package hq4x

import (
	"image"
	"image/color"
)

// Scale4x enlarges src to four times its width and height with HQ4x.
// It uses the standard HQX thresholds and does not wrap at image edges.
// The returned image starts at (0, 0) and uses non-premultiplied alpha.
func Scale4x(src image.Image) *image.NRGBA {
	bounds := src.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	maxInt := int(^uint(0) >> 1)
	if width < 0 || height < 0 || width > maxInt/4 || height > maxInt/4 || (width != 0 && height > maxInt/64/width) {
		panic("hqx: invalid image dimensions")
	}
	in := make([]int, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := color.NRGBAModel.Convert(src.At(bounds.Min.X+x, bounds.Min.Y+y)).(color.NRGBA)
			in[y*width+x] = int(uint32(c.A)<<24 | uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B))
		}
	}
	out := make([]int, len(in)*16)
	Hq4x_32_rb(in, out, width, height, 48, 7, 6, 0, false, false)
	dst := image.NewNRGBA(image.Rect(0, 0, width*4, height*4))
	for i, pixel := range out {
		c := uint32(pixel)
		dst.Pix[i*4] = byte(c >> 16)
		dst.Pix[i*4+1] = byte(c >> 8)
		dst.Pix[i*4+2] = byte(c)
		dst.Pix[i*4+3] = byte(c >> 24)
	}
	return dst
}
