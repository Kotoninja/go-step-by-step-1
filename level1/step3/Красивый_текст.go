package main

import (
	"fmt"
	"strings"
)

func main() {
	var number_of_count int
	fmt.Scanln(&number_of_count)

	var banned_word string
	fmt.Scanln(&banned_word)

	var count int
	for i := 0; i < number_of_count; i++ {
		var word string
		fmt.Scan(&word)

		if strings.ToLower(banned_word) == strings.ToLower(word) {
			count++
		}

	}
	fmt.Println(count)
}
