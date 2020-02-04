package main

import (
	"image"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"

	"image/color"
	"image/jpeg"
	_ "image/jpeg"
)

type Changeable interface {
	Set(x, y int, c color.Color)
}

func main() {
	// Decode the JPEG data.
	reader, err := os.Open("resources/Arches.jpg")
	if err != nil {
		panic(err.Error())
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)

	if err != nil {
		panic(err.Error())
	}

	bounds := m.Bounds()
	img := image.NewRGBA(bounds)

	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			var ur uint16 = 255 - uint16(r)
			var ug uint16 = 255 - uint16(g)
			var ub uint16 = 255 - uint16(b)
			c := color.RGBA64{ur, ug, ub, uint16(a)}

			img.Set(x, y, c)
		}
	}

	var opt jpeg.Options

	opt.Quality = 100
	jpeg.Encode(os.Stdout, img, &opt)
}
