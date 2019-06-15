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
			t.Errorf("got %.2f want %.2f.", got, want)
		}
	}

	t.Run("Calculate rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		checkArea(t, &rectangle, 50.0)

	})

	t.Run("Calculate area with negatives", func(t *testing.T) {
		rectangle := Rectangle{-10.0, 5.0}
		checkArea(t, &rectangle, 0)
	})

	t.Run("Calculate circle area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, &circle, 314.1592653589793)
	})
}
