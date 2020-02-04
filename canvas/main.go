// canvas processes images with different filters.
// usage:
// 		canvas ``input file'' -f ``filter'' ``output file''
// 		available filters: -[b]righten, -[g]rayscale, -[i]nvert
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

var filterArg = flag.String(
	"f",
	"",
	"filter options: [b]righten, [g]rayscale, [i]nvert",
)

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

	f, ok := filters[*filterArg]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	modify(os.Stdout, img, f)
}

// Return a copy of the image modified by the input filter
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

// Returns a JPEG as an Image
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
