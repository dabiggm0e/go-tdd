package shapes

import "testing"

func assertCorrectRectangleSize(t *testing.T, got, want float64, r Rectangle) {
	if got != want {
		t.Errorf("got %.2f want %.2f. %.2f %.2f", got, want, r.Height, r.Width)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("Calculate perimeter", func(t *testing.T) {

		rectangle := Rectangle{10.5, 12.5}
		got := Perimeter(rectangle)
		want := float64(46.0)

		assertCorrectRectangleSize(t, got, want, rectangle)
	})

	t.Run("Calculate perimeter with negatives", func(t *testing.T) {

		rectangle := Rectangle{-10.5, 12.5}
		got := Perimeter(rectangle)
		want := float64(0.0)

		assertCorrectRectangleSize(t, got, want, rectangle)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f.", shape, got, want)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 10.0, Height: 5.0}, hasArea: 50.0},
		{name: "Rectangle", shape: &Rectangle{Width: -10.0, Height: 5.0}, hasArea: 0},
		{name: "Circle", shape: &Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{Base: 10, Height: 4}, hasArea: 20},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})

	}

}
