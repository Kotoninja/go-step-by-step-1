package main

import (
	"fmt"
)

func main() {
	var number_of_pupil int
	fmt.Scanln(&number_of_pupil)

	var discount float64
	fmt.Scanln(&discount)
	discount /= 100

	var total_sum float64

	for i := 0; i < number_of_pupil; i++ {
		var product_cost float64
		fmt.Scanln(&product_cost)

		product_cost -= product_cost * discount

		total_sum += product_cost
	}

	fmt.Println(total_sum)
}
