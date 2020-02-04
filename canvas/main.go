package main

import (
	"image"
	"io"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"

	"image/jpeg"
	_ "image/jpeg"

	"github.com/RosyGraph/canvas/filter"
)

func main() {
	// TODO: add flag functionality
	// TODO: add gif/png functionality
	img := decodeJPEG("resources/Arches.jpg")
	modifyAll(os.Stdout, img, filter.Brighten)
}

func modifyAll(writer io.Writer, m image.Image, f filter.Filter) {
	bounds := m.Bounds()
	img := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := m.At(x, y)

			cc := f(c)
			img.Set(x, y, cc)
		}
	}
	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(writer, img, &opt)
}

func decodeJPEG(f string) image.Image {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	m, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}
	return m
}
