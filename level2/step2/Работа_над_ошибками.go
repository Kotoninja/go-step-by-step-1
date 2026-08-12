package main

import "fmt"

func myFunc() int {
	// Задаём переменную
	a := 1
	return a
}

func main() {
	// Вызываем переменную
	a := myFunc()

	// Пробуем вывести переменную из функции myFunc в консоль
	fmt.Println(a)
}
