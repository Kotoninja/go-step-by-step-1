package main

import (
	"fmt"
)

func main() {
	var first string
	var second string

	fmt.Scanln(&first)
	fmt.Scanln(&second)

	switch {
	case first == second:
		fmt.Println("Ничья")
	case (first == "бумага" && second == "камень") || (first == "ножницы" && second == "бумага"):
		fmt.Println("Первый игрок победил")
	default:
		fmt.Println("Второй игрок победил")
	}

}
