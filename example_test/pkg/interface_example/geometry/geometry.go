package geometry

import "math"

type Triangle struct {
	size int
}

func (t *Triangle) DoubleSize() {
	t.size *= 2
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) GetSize() (size int) {
	return t.size
}

func (t *Triangle) Perimeter() int {
	return t.size * 3
}

type Square struct {
	Size float64
}

func (s Square) Area() float64 {
	return s.Size * s.Size
}

func (s Square) Perimeter() float64 {
	return s.Size * 4
}

type ColoredTriangle struct {
	Triangle
	Color string
}

func (t *ColoredTriangle) Perimeter() int {
	return t.size * 3 * 2
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
