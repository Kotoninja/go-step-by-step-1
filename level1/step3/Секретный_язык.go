package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scanln(&word)

	var total_word string

	for _, letter := range word {
		if letter != 'а' && letter != 'о' {
			total_word += string(letter)
		}
	}

	fmt.Println(total_word)
}
