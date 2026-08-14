package main

import (
	"fmt"
)

func SumOfValuesInMap(m map[int]int) int {
	var result int = 0

	for _, v := range m {
		result += v
	}

	return result
}

func main() {
	m := map[int]int{10: 37, 3: 19}
	fmt.Println(SumOfValuesInMap(m))
}
