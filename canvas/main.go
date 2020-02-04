package main

import (
	"image"
	"io"
	"os"

	// _ "image/gif"
	// _ "image/png"

	"image/color"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/RosyGraph/canvas/filter"
)

func main() {
	// TODO: add flag functionality
	// TODO: add gif/png functionality

	img := decodeJPEG("resources/Arches.jpg")
	brighten := func(c color.Color) color.Color {
		return filter.Brighten(c, 2)
	}
	modifyAll(os.Stdout, img, brighten)
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
