// xorsearch.go
// description: Find a missing number in a sequence
// details:
// Given an array A containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array. - [xor](https://en.wikipedia.org/wiki/Exclusive_or)
// author(s) [red_byte](https://github.com/i-redbyte)
// see xorsearch_test.go

package binary

func XorSearchMissingNumber(a []int) int {
	n := len(a)
	result := len(a)
	for i := 0; i < n; i++ {
		result ^= i ^ a[i]
	}
	return result
}
