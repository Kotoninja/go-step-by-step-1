package main

import (
	"fmt"
	"math"
)

func FindMaxKey(m map[int]int) int {
	var result float64 = math.Inf(-1)

	for k := range m {
		result = math.Max(result, float64(k))
	}

	return int(result)
}

func main() {
	m := map[int]int{10: 37, 3: 19}
	fmt.Println(FindMaxKey(m))
}
