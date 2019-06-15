//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration
package iteration

// Repeat returns the letter in input repeated by the count parameter
func Repeat(ch string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += ch
	}
	return repeated
}
