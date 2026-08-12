package main

import (
	"fmt"
)

func main() {
	var first float64
	var second float64
	var third float64

	fmt.Scan(&first)
	fmt.Scan(&second)
	fmt.Scan(&third)

	if first == second && second == third {
		fmt.Println("Максимальное равенство")
		return
	}
	fmt.Println("Не равны")
}
