package main

import "fmt"

func GoOrNot(phrase string) {
	if phrase == "Go" {
		fmt.Println("Go!")
		return
	}

	fmt.Println("Я знаю только Go!")
}