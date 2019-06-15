package shapes

//Perimeter returns the perimeter of a rectangle given its height and width
func Perimeter(height, width float64) float64 {

	if height*width < 0 {
		return 0
	}
	perimeter := 2.0 * (height + width)
	return perimeter
}
