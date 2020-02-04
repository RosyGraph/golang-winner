package main

import (
	"flag"
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

var farg = flag.String(
	"filter",
	"",
	"filter options: [b]righten, [g]rayscale, [i]nvert",
)

// TODO: comment
func main() {
	// TODO: add gif/png functionality
	img := decodeJPEG("resources/Arches.jpg")

	brighten := func(c color.Color) color.Color {
		return filter.Brighten(c, 2)
	}

	filters := map[string]filter.Filter{
		"brighten":  brighten,
		"b":         brighten,
		"grayscale": filter.Grayscale,
		"g":         filter.Grayscale,
		"invert":    filter.Invert,
		"i":         filter.Invert,
	}

	flag.Parse()

	var f filter.Filter

	f, ok := filters[*farg]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	modify(os.Stdout, img, f)
}

// TODO: comment
func modify(writer io.Writer, m image.Image, f filter.Filter) {
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

// TODO: comment
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
