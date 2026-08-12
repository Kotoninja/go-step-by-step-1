package main

import (
	"fmt"
	"math"
)

func main() {
	var number int

	fmt.Scanln(&number)

	if number == 0 {
		fmt.Println("Число 0")
	} else if int(math.Log10(float64(number))) == 0 {
		fmt.Println("Число однозначное")
	} else if number%2 == 0 {
		fmt.Println("Число чётное")
	} else if number > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число красивое")
	}

}
