package main

import (
	"fmt"
	"math"
)

func main() {
	var first int
	var second int
	var third int

	fmt.Scanln(&first)
	fmt.Scanln(&second)
	fmt.Scanln(&third)

	fmt.Println(math.Min(float64(third), math.Min(float64(first), float64(second))))
}
