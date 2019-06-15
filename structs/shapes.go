package shapes

type Rectangle struct {
	Height float64
	Width  float64
}

//Perimeter returns the perimeter of a rectangle given its height and width
func Perimeter(r Rectangle) float64 {

	if r.Height*r.Width < 0 {
		return 0
	}
	perimeter := 2.0 * (r.Height + r.Width)
	return perimeter
}

func Area(rectangle Rectangle) float64 {
	area := rectangle.Height * rectangle.Width
	if area < 0 {
		return 0
	}
	return area
}
