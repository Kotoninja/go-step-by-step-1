package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println(-1)
		return
	}

	result := math.Sqrt(number)

	fmt.Printf("%.3f\n", result)
}
