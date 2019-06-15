package shapes

import "testing"

func assertCorrectPerimeter(t *testing.T, got, want, height, width float64) {
	if got != want {
		t.Errorf("got %f want %f. %f %f", got, want, height, width)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("Calculate perimeter", func(t *testing.T) {
		height, width := float64(10.5), float64(12.5)
		got := Perimeter(height, width)
		want := float64(46.0)

		assertCorrectPerimeter(t, got, want, height, width)
	})

	t.Run("Calculate perimeter with negatives", func(t *testing.T) {
		height, width := float64(-10.5), float64(12.5)
		got := Perimeter(height, width)
		want := float64(0.0)

		assertCorrectPerimeter(t, got, want, height, width)
	})
}
