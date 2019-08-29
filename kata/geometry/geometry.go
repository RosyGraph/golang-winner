package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Perimeter() float64 {
	return 0.0
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2.0
}

type Square struct {
	side float64
}

func (s Square) Perimeter() float64 {
	return s.side * 4.0
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Perimeter() float64 {
	return r.length*2.0 + r.width*2.0
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
