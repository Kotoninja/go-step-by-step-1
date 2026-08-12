package main

import (
	"fmt"
)

func main() {

	for {
		var phrase string
		fmt.Scanln(&phrase)

		if phrase == "да" || phrase == "нет" || phrase == "чёрный" || phrase == "белый" {
			fmt.Println("Поражение")
			break
		}
		fmt.Println("Игра продолжается")
	}
}
