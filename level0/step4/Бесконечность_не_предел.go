package main

import (
	"fmt"
	"math"
)

func main() {
	var first float64
	var second float64

	fmt.Scanln(&first)
	fmt.Scanln(&second)

	fmt.Println(math.Floor(math.Max(first, second)))
}
