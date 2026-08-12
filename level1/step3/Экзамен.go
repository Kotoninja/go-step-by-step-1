package main

import (
	"fmt"
)

func main() {
	var number_of_pupil int

	fmt.Scanln(&number_of_pupil)

	for i := 0; i < number_of_pupil; i++ {
		var mark int

		fmt.Scanln(&mark)

		switch {
		case mark >= 90 && mark <= 100:
			fmt.Println(5)
		case mark >= 75 && mark <= 89:
			fmt.Println(4)
		case mark >= 50 && mark <= 74:
			fmt.Println(3)
		case mark >= 0 && mark <= 49:
			fmt.Println(2)
		default:
			fmt.Println("Неверный балл")
		}
	}
}
