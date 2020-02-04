package filter

import (
	"image/color"
)

type Filter func(c color.Color) color.Color

func Brighten(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	cc := color.RGBA64{
		R: safe(r * 2),
		G: safe(g * 2),
		B: safe(b * 2),
		A: uint16(a),
	}
	return cc
}

func Grayscale(c color.Color) color.Color {
	return color.Gray16Model.Convert(c)
}

func Invert(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	cc := color.RGBA64{
		R: 255 - uint16(r),
		G: 255 - uint16(g),
		B: 255 - uint16(b),
		A: uint16(a),
	}
	return cc
}

func safe(n uint32) uint16 {
	switch {
	case n > 65535:
		n = 65535
	case n < 0:
		n = 0
	}
	return uint16(n)
}
