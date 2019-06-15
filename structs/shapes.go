//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces
package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

//Perimeter returns the perimeter of a rectangle given its height and width
func Perimeter(r Rectangle) float64 {

	if r.Height*r.Width < 0 {
		return 0
	}
	perimeter := 2.0 * (r.Height + r.Width)
	return perimeter
}

func (rectangle *Rectangle) Area() float64 {
	area := rectangle.Height * rectangle.Width
	if area < 0 {
		return 0
	}
	return area
}

func (circle *Circle) Area() float64 {
	//return math.Pow(circle.Radius, 2) * math.Pi
	return circle.Radius * circle.Radius * math.Pi
}

func (t *Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
