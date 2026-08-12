package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var input_time_one string
	var input_time_two string

	fmt.Scanln(&input_time_one)
	fmt.Scanln(&input_time_two)

	time_one, _ := time.Parse("2006-01-02", input_time_one)
	time_two, _ := time.Parse("2006-01-02", input_time_two)

	duration := int(math.Abs(float64(time_two.Year() - time_one.Year())))

	s := fmt.Sprintf("%v year ago", duration)
	fmt.Println(s)
}
