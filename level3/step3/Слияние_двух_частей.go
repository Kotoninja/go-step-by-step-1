package main

func Mix(nums []int) []int {
	var result []int

	for i := 0; i < len(nums)/2; i++ {
		result = append(result, nums[i])
		result = append(result, nums[i+len(nums)/2])
	}
	return result
}
