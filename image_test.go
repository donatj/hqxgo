package hq4x

import (
	"image"
	"image/color"
	"testing"
)

func TestScale4x(t *testing.T) {
	src := image.NewNRGBA(image.Rect(5, 7, 6, 8))
	want := color.NRGBA{R: 18, G: 52, B: 86, A: 127}
	src.SetNRGBA(5, 7, want)
	got := Scale4x(src)
	if got.Bounds() != image.Rect(0, 0, 4, 4) {
		t.Fatalf("bounds: %v", got.Bounds())
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if c := got.NRGBAAt(x, y); c != want {
				t.Fatalf("pixel %d,%d: got %v, want %v", x, y, c, want)
			}
		}
	}
}

func TestScale4xEmpty(t *testing.T) {
	got := Scale4x(image.NewNRGBA(image.Rect(0, 0, 0, 0)))
	if got.Bounds() != image.ZR {
		t.Fatalf("bounds: %v", got.Bounds())
	}
}
