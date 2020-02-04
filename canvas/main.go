// canvas processes images with different filters.
// usage:
// 		canvas ``input file'' -f ``filter'' ``output file''
// 		available filters: -[b]righten, -[g]rayscale, -[i]nvert
package main

import (
	"bufio"
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
	"golang.org/x/crypto/openpgp/errors"
)

func main() {
	// TODO: add gif/png functionality
	var (
		flagIn     = flag.String("in", "", "Input filename")
		flagFilter = flag.String("f", "", "Filter name")
		flagOut    = flag.String("out", "", "Output filename")
		f          filter.Filter
		m          image.Image
		w          io.Writer
		err        error
	)

	flag.Parse()

	if m, err = decodeJPEG(*flagIn); err != nil {
		panic(err)
	}

	if f, err = parseFilterArg(*flagFilter); err != nil {
		panic(err)
	}

	if w, err = parseOutArg(*flagOut); err != nil {
		panic(err)
	}

	modify(w, m, f)
}

func parseOutArg(s string) (io.Writer, error) {
	var w io.Writer
	if s == "" {
		w = os.Stdout
	} else {
		file, err := os.Create(s)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		w = bufio.NewWriter(file)
	}
	return w, nil
}

// Parse the input string pointer into a filter option
func parseFilterArg(s string) (filter.Filter, error) {
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
	f, ok := filters[s]
	if !ok {
		return nil, errors.InvalidArgumentError(
			"Usage: -f [b|brighten|g|grayscale|i|invert]")
	}
	return f, nil
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
func decodeJPEG(f string) (image.Image, error) {
	r, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	m, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return m, nil
}
