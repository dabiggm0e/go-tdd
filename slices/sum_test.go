package slices

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertSum := func(t *testing.T, numbers []int, sum, expected int) {
		if sum != expected {
			t.Errorf("%v. got '%d' want '%d'", numbers, sum, expected)
		}
	}

	t.Run("Sum of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6
		assertSum(t, numbers, sum, expected)
	})
}
