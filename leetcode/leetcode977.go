package leetcode

import (
	"sort"
)

func sortedSquares(A []int) []int {
	length := len(A)
	if length <= 1 {
		return A
	}
	for i := 0; i < length; i++ {
		A[i] = A[i]*A[i]
	}
	sort.Ints(A)
	return A
}

/*不需要排序的方案*/
func sortedSquares2(A []int) []int {
	length := len(A)
	if length <= 1 {
		return A
	}
	for i := 0; i < length; i++ {
		A[i] = A[i]*A[i]
	}
	sort.Ints(A)
	return A
}