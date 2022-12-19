package geom

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return perimeter
}

func (c Circle) Perimeter() (perimeter float64) {
	perimeter = 2 * math.Pi * c.Radius
	return perimeter
}

func (tr Triangle) Perimeter() (perimeter float64) {
	perimeter = tr.Base*2 + tr.Height*2
	return perimeter
}

func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return area
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return area
}

func (tr Triangle) Area() (area float64) {
	area = 0.5 * tr.Base * tr.Height
	return area
}
