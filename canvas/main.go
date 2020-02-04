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

	"image/color"
	"image/jpeg"
	_ "image/jpeg"
)

type filter func(c color.Color) color.Color

func main() {
	// TODO: add flag functionality
	img := decodeJPEG("resources/Arches.jpg")
	// reverse(os.Stdout, img)
	modifyAll(os.Stdout, img, grayscale)
}

func modifyAll(writer io.Writer, m image.Image, f filter) {
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

func grayscale(c color.Color) color.Color {
	return color.Gray16Model.Convert(c)
}

func invert(c color.Color) color.Color {
	r, g, b, a := c.RGBA()

	cc := color.RGBA64{
		R: 255 - uint16(r),
		G: 255 - uint16(g),
		B: 255 - uint16(b),
		A: uint16(a),
	}

	return cc
}

func reverse(writer io.Writer, m image.Image) {
	bounds := m.Bounds()
	img := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := m.At(x, y)

			// Invert each color value
			img.Set(x, y, invert(c))
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
