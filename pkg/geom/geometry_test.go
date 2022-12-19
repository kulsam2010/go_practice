package geom

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	testPerimeters := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 2, Height: 3}, want: 10.0},
		{shape: Circle{Radius: 10}, want: 2 * math.Pi * 10},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range testPerimeters {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("%v got %.2f, wanted %.2f", tt, got, tt.want)
		}
	}

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%v got %.2f, wanted %.2f", tt, got, tt.want)
		}
	}

}
