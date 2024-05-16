package min_distinct

import (
	"sort"
)

func solution(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	ans := 0

	for i, a := range A {
		ans += abs(i + 1 - a)
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
