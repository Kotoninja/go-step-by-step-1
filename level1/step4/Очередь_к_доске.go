package main

import (
	"fmt"
)

func main() {
	var one_name string = "-"
	var one_busy bool = false

	var two_name string = "-"
	var two_busy bool = false

	var three_name string = "-"
	var three_busy bool = false

	var four_name string = "-"
	var four_busy bool = false

	var five_name string = "-"
	var five_busy bool = false

	var place int = 0

	MainLoop:
		for {
			var command string
			fmt.Scan(&command)

			switch command {
			case "":
				continue
			case "очередь":
				fmt.Printf("%d. %v\n", 1, one_name)
				fmt.Printf("%d. %v\n", 2, two_name)
				fmt.Printf("%d. %v\n", 3, three_name)
				fmt.Printf("%d. %v\n", 4, four_name)
				fmt.Printf("%d. %v\n", 5, five_name)
				continue
			case "количество":
				fmt.Printf("Осталось свободных мест: %d\n", 5-place)
				fmt.Printf("Всего человек в очереди: %d\n", place)
				continue
			case "конец":
				fmt.Printf("%d. %v\n", 1, one_name)
				fmt.Printf("%d. %v\n", 2, two_name)
				fmt.Printf("%d. %v\n", 3, three_name)
				fmt.Printf("%d. %v\n", 4, four_name)
				fmt.Printf("%d. %v\n", 5, five_name)
				break MainLoop
			}

			var number int
			fmt.Scan(&number)

			if number < 1 || number > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", number)
				continue
			}

			if place >= 5 {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}

			if number == 1 {
				if one_name == "-" {
					one_name = command
					place++
					one_busy = true
					continue
				}
				if one_busy && one_name != "-" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}

			if number == 2 {
				if two_name == "-" {
					two_name = command
					place++
					two_busy = true
					continue
				}
				if two_busy && two_name != "-" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}

			if number == 3 {
				if three_name == "-" {
					three_name = command
					place++
					three_busy = true
					continue
				}
				if three_busy && three_name != "-" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}

			if number == 4 {
				if four_name == "-" {
					four_name = command
					place++
					four_busy = true
					continue
				}
				if four_busy && four_name != "-" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}

			if number == 5 {
				if five_name == "-" {
					five_name = command
					place++
					five_busy = true
					continue
				}
				if five_busy && five_name != "-" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
				continue
			}
		}
	}
