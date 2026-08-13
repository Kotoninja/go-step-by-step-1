package main

func FindMaxMinInArray(array [10]int) (int, int) {
	var maxDigit int = array[0]
	var minDigit int = array[0]

	for _, digit := range array {
		maxDigit = max(maxDigit, digit)
		minDigit = min(minDigit, digit)
	}

	return maxDigit, minDigit
}
