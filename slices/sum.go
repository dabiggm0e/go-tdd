//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices
package slices

//Sum returns the sum of the slice of integers in input
func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

//SumAll returns a slice summing each slice in the input
func SumAll(slices ...[]int) []int {
	var sumAll []int

	for _, slice := range slices {
		sumAll = append(sumAll, Sum(slice))
	}

	return sumAll
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(slice[1:]))

	}

	return sums
}
