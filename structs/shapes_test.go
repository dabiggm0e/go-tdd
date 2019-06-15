package shapes

import "testing"

func assertCorrectShapeSize(t *testing.T, got, want, height, width float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f. %.2f %.2f", got, want, height, width)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("Calculate perimeter", func(t *testing.T) {
		height, width := float64(10.5), float64(12.5)
		got := Perimeter(height, width)
		want := float64(46.0)

		assertCorrectShapeSize(t, got, want, height, width)
	})

	t.Run("Calculate perimeter with negatives", func(t *testing.T) {
		height, width := float64(-10.5), float64(12.5)
		got := Perimeter(height, width)
		want := float64(0.0)

		assertCorrectShapeSize(t, got, want, height, width)
	})
}

func TestArea(t *testing.T) {
	t.Run("Calculate area", func(t *testing.T) {
		height := 10.0
		width := 5.0
		got := Area(height, width)
		want := 50.0

		assertCorrectShapeSize(t, got, want, height, width)
	})

	t.Run("Calculate area with negatives", func(t *testing.T) {
		height := -10.0
		width := 5.0
		got := Area(height, width)
		want := 0.0

		assertCorrectShapeSize(t, got, want, height, width)
	})
}
