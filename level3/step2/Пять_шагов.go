package main

func FiveSteps(array [5]int) [5]int {
	var answer [5]int

	for i := len(array) - 1; i >= 0; i-- {
		answer[len(array)-1-i] = array[i]
	}

	return answer
}
