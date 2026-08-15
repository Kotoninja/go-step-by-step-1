package main

import (
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)
	result := append(left, right...)
	sort.Ints(result)
	return result
}
