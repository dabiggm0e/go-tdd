package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertSum := func(t *testing.T, numbers []int, sum, want int) {
		if sum != want {
			t.Errorf("%v. got '%d' want '%d'", numbers, sum, want)
		}
	}

	t.Run("Sum of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6
		assertSum(t, numbers, sum, want)
	})
}

func assertSliceSum(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{5, 6, 7, 8}
	slice3 := []int{}

	want := []int{6, 26, 0}
	got := SumAll(slice1, slice2, slice3)

	assertSliceSum(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum tails of slices", func(t *testing.T) {
		slice1 := []int{5}
		slice2 := []int{1, 2, 3, 4}
		slice3 := []int{6, 7}

		want := []int{0, 9, 7}
		got := SumAllTails(slice1, slice2, slice3)

		assertSliceSum(t, got, want)
	})

	t.Run("Sum tails of slices safely with empty slice", func(t *testing.T) {
		slice1 := []int{5}
		slice2 := []int{1, 2, 3, 4}
		slice3 := []int{}

		want := []int{0, 9, 0}
		got := SumAllTails(slice1, slice2, slice3)

		assertSliceSum(t, got, want)
	})
}
