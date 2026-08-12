package main

import (
	"fmt"
)

func main() {
	var sign string
	var weather int

	fmt.Scan(&sign)
	fmt.Scan(&weather)
    
	if sign == "+" && weather > 20 {
		fmt.Println("Стоит надеть майку и шорты")
		return
	}

	if sign == "+" && (weather >= 10 && weather <= 20) {
		fmt.Println("Стоит надеть штаны и кофту")
		return
	}

	if (weather == 0) || ((sign == "-" && (weather > 0 && weather <= 5)) || (sign == "+" && (weather > 0 && weather <= 9))) {
		fmt.Println("Стоит надеть куртку")
		return
	}

	fmt.Println("Стоит надеть зимнюю куртку")
}
