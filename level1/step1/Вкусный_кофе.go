package main

import (
	"fmt"
)

func main() {
	var rub int
	var cop int
	var cost int

	fmt.Scan(&rub)
	fmt.Scan(&cop)
	fmt.Scan(&cost)

	have_money := rub + cop/100

	if have_money >= cost {
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}

	fmt.Println("Стоит подкопить")
}
