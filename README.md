# HQX Go

[![CI](https://github.com/donatj/hqxgo/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/hqxgo/actions/workflows/ci.yml)

HQX Go scales pixel art using the HQ4x filter. It enlarges an image by four
in each dimension. It is a port of [hqx-java](https://github.com/Arcnor/hqx-java).
hqx-java is a port of hqxSharp, which is a port of the original hqx project.

## Command line

Install the command with `go install github.com/donatj/hqxgo/cmd/hqxgo@latest`.
Run it with an input path and an output path:

```sh
hqxgo input.png output.png
```

The command accepts PNG, JPEG, and GIF images. It writes PNG output. For an
animated GIF, it uses the first frame. Existing output files are overwritten.

## Library usage

```go
import (
    "image"

    hqx "github.com/donatj/hqxgo"
)

func enlarge(src image.Image) *image.NRGBA {
    return hqx.Scale4x(src)
}
```

`Scale4x` accepts any Go image. It returns an image with non-premultiplied
alpha and bounds starting at `(0, 0)`. It uses the standard thresholds and
repeats edge pixels at the image boundary.

For custom thresholds or wrapping, use `Hq4x_32_rb`. Its input and output
slices contain packed, non-premultiplied alpha, red, green, blue (ARGB) pixels.
The destination needs 16 entries per source pixel. The slices must not overlap.
The function panics on invalid dimensions, thresholds, or buffer sizes.

Alpha comparisons use the full unsigned range from 0 to 255. This fixes the
signed alpha comparison in hqx-java. Color conversion happens on demand. It
does not require lookup table initialization.

Only the 4× filter is implemented. The 2× and 3× filters are not included.

## License

The port is derived from hqx-java. It is available under the GNU Lesser General
Public License, version 3 or later. See [LGPL-3.0.txt](LGPL-3.0.txt) and
[GPL-3.0.txt](GPL-3.0.txt).
