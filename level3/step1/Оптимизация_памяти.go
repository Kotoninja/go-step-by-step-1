package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	answer := first + second

	return fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", answer, len(answer), utf8.RuneCountInString(answer))
}

func main() {
	fmt.Println(CountLengthAndBytes("first", "second"))
}
