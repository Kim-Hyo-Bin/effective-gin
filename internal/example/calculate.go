package example

import "math"

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

type Circle struct {
	Radius float64
}

// Area and Perimeter methods for rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Area and Perimeter methods for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
