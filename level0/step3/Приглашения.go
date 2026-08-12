package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)

	var room int
	fmt.Scanln(&room)

	var password int
	fmt.Scanln(&password)

	var hours float64
	fmt.Scanln(&hours)

	fmt.Println(fmt.Sprintf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.", name, room, hours, password))
}
