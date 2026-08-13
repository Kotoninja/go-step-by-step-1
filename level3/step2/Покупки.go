package main

func SumOfArray(array [6]int) int {
	var sum int

	for _, number := range array {
		sum += number
	}

	return sum
}
