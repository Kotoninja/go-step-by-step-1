package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i, item := range array {
		if i <= 6 {
			fmt.Printf("%d я уже сделал: %s\n", i+1, item)
		} else {
			fmt.Printf("%d не успел сделать: %s\n", i+1, item)
		}
	}
}

func main() {
	input := [9]string{
		"проснуться",
		"позавтракать",
		"сходить в школу",
		"пообедать",
		"погулять с друзьями",
		"сделать домашнюю работу",
		"попрограммировать на Go",
		"поужинать",
		"лечь спать",
	}
	PrettyArrayOutput(input)
}
