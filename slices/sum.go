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
