package main

func Join(nums1, nums2 []int) []int {
	result := make([]int, (len(nums1) + len(nums2)))

	copy(result, nums1)
	copy(result[len(nums1):], nums2)
	// for i, value := range nums1 {
	// 	result[i] = value
	// }

	// for i, value := range nums2 {
	// 	result[len(nums2)+i] = value
	// }

	return result
}
