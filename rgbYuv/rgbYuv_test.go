package rgbYuv

import "testing"

func TestGetYuvPrecision(t *testing.T) {
	if got := GetYuv(0x004818) >> 16; got != 44 {
		t.Fatalf("luminance = %d, want 44", got)
	}
	withAlpha := uint32(0xff004818)
	if GetYuv(int(withAlpha)) != GetYuv(0x004818) {
		t.Fatal("alpha changed YUV result")
	}
}
