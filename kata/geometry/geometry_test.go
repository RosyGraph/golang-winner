package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "area of a square", shape: Square{3}, want: 9.0},
		{name: "area of a rectangle", shape: Rectangle{3, 4}, want: 12.0},
		{name: "area of a circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "area of a triangle", shape: Triangle{2, 3}, want: 3},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assertArea(t, c.shape, c.want)
		})
	}
}

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "perimeter of a square", shape: Square{3}, want: 12.0},
		{name: "perimeter of a rectangle", shape: Rectangle{3, 4}, want: 14.0},
		{name: "perimeter of a circle", shape: Circle{1}, want: 2 * math.Pi},
		// {name: "perimeter of a triangle", shape: Triangle{2, 3}, want: 3},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assertPerimeter(t, c.shape, c.want)
		})
	}
}

func assertPerimeter(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
