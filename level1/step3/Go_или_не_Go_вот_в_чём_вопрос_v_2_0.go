package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		var input string

		fmt.Scanln(&input)

		if input == "Go" {
			fmt.Println("Go!")
			continue
		}

		fmt.Println("Я знаю только Go!")
	}
}
