package main

import "unicode"

func hasMinimumLength(password string, minLength int) bool {
	return len(password) >= minLength
}

func hasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}

	return false
}

func hasLowerCase(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}

	return false
}

func ratePassword(password string) string {
	var passed_checks int

	if hasLowerCase(password) {
		passed_checks++
	}
	if hasUpper(password) {
		passed_checks++
	}
	if hasMinimumLength(password, 8) {
		passed_checks++
	}

	switch passed_checks {
	case 3:
		return "Сложный пароль"
	case 2:
		return "Средний пароль"
	case 1:
		return "Слабый пароль"
	default:
		return "Пароль недостаточно безопасен, придумайте новый"
	}

}
