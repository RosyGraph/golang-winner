package filter

import (
	"image/color"
)

type Filter func(c color.Color) color.Color

// Return a copy of the color brigthened by a factor of f
func Brighten(c color.Color, f uint32) color.Color {
	// NOTE: Brighten does not implement Filter and must be
	// wrapped to use modify() in the main function
	r, g, b, a := c.RGBA()
	cc := color.RGBA64{
		R: safe(r * f),
		G: safe(g * f),
		B: safe(b * f),
		A: uint16(a),
	}
	return cc
}

// Return a copy of the color converted to 16bit grayscale
func Grayscale(c color.Color) color.Color {
	return color.Gray16Model.Convert(c)
}

// Return a copy of the color inverted
func Invert(c color.Color) color.Color {
	// FIXME: safe colors
	r, g, b, a := c.RGBA()
	cc := color.RGBA64{
		R: 255 - uint16(r),
		G: 255 - uint16(g),
		B: 255 - uint16(b),
		A: uint16(a),
	}
	return cc
}

// Convert a potentially unsafe uint32 color value into a
// safe uint16 color value
func safe(n uint32) uint16 {
	switch {
	case n > 65535:
		n = 65535
	case n < 0:
		n = 0
	}
	return uint16(n)
}
