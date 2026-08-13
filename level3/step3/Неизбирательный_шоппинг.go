package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("N <= 0")
	} else if nums == nil {
		return nil, errors.New("not enough return values")
	}

	var result []int

	for _, v := range nums {
		if n <= 0 {
			break
		}
		if v < limit {
			result = append(result, v)
		}
	}
	return result, nil
}
