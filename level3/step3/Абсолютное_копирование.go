package main

func SliceCopy(nums []int) []int {
	return make([]int, cap(nums), cap(nums))
}
