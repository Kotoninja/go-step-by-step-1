package main

import (
	"fmt"
	"math"
)

func main() {
	var first float64
	var second float64

	fmt.Scan(&first)
	fmt.Scan(&second)

	if first == second {
		fmt.Println("Числа равны")
		return 
	}

	if greater := math.Max(first, second); greater == first {
		fmt.Println("Первое число больше второго")
		return
	}

	if greater := math.Max(first, second); greater == second {
		fmt.Println("Второе число больше первого")
		return
	}

}
