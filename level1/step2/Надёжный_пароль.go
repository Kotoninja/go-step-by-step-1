package main

import (
	"fmt"
)

func main() {
	var first_password string
	var second_password string

	fmt.Scanln(&first_password)
	fmt.Scanln(&second_password)

	switch {
	case len(first_password) >= 8 && len(second_password) >= 8:
		fmt.Println("Оба пароля надёжные")
	case len(first_password) < 8 && len(second_password) < 8:
		fmt.Println("Оба пароля ненадёжные")
	case len(first_password) >= 8 && len(second_password) < 8:
		fmt.Println("Только первый пароль надёжный")
	case len(first_password) < 8 && len(second_password) >= 8:
		fmt.Println("Только второй пароль надёжный")
	}

}
