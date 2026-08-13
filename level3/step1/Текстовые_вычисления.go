package main

import "fmt"

func switchToRune(char rune) string {
	switch char {
	case '0':
		return "ноль"
	case '1':
		return "один"
	case '2':
		return "два"
	case '3':
		return "три"
	case '4':
		return "четыре"
	case '5':
		return "пять"
	case '6':
		return "шесть"
	case '7':
		return "семь"
	case '8':
		return "восемь"
	case '9':
		return "девять"
	case '+':
		return "плюс"
	case '-':
		return "минус"
	case '/':
		return "делить на"
	case '*':
		return "умножить на"
	default:
		return string(char)
	}
}

func NumbersToLetters(input string) string {
	var answer string

	for _, char := range input {
		answer = answer + switchToRune(char)
	}

	return answer
}

func main() {
	fmt.Println(NumbersToLetters("(1 + 2) * 3 / 8"))
}
