package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()
	for _, format := range []string{"png", "jpeg"} {
		t.Run(format, func(t *testing.T) {
			input := filepath.Join(dir, "input."+format)
			output := filepath.Join(dir, "output-"+format+".png")
			f, err := os.Create(input)
			if err != nil {
				t.Fatal(err)
			}
			src := image.NewNRGBA(image.Rect(0, 0, 2, 1))
			src.SetNRGBA(0, 0, color.NRGBA{R: 255, A: 255})
			src.SetNRGBA(1, 0, color.NRGBA{B: 255, A: 255})
			if format == "png" {
				err = png.Encode(f, src)
			} else {
				err = jpeg.Encode(f, src, nil)
			}
			if err != nil {
				t.Fatal(err)
			}
			if err = f.Close(); err != nil {
				t.Fatal(err)
			}
			if err = run([]string{input, output}); err != nil {
				t.Fatal(err)
			}
			f, err = os.Open(output)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			got, err := png.Decode(f)
			if err != nil {
				t.Fatal(err)
			}
			if got.Bounds() != image.Rect(0, 0, 8, 4) {
				t.Fatalf("bounds = %v", got.Bounds())
			}
			if format == "png" {
				want := color.NRGBAModel.Convert(src.At(0, 0))
				if pixel := color.NRGBAModel.Convert(got.At(0, 0)); pixel != want {
					t.Fatalf("pixel = %v, want %v", pixel, want)
				}
			}
		})
	}
}

func TestRunErrors(t *testing.T) {
	if err := run(nil); err == nil {
		t.Fatal("missing arguments accepted")
	}
	dir := t.TempDir()
	input, output := filepath.Join(dir, "invalid"), filepath.Join(dir, "output.png")
	if err := os.WriteFile(input, []byte("not an image"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{input, output}); err == nil {
		t.Fatal("invalid image accepted")
	}
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		t.Fatalf("output created for invalid input: %v", err)
	}
}
