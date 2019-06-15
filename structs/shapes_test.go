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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{10.0, 5.0}, 50.0},
		{&Rectangle{-10.0, 5.0}, 0},
		{&Circle{10}, 314.1592653589793},
		{&Triangle{10, 4}, 20},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}

}
