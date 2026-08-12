package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var date string
	var first_name string
	var second_name string
	var surname string
	var cost_one float64
	var cost_two float64
	var cost_third float64

	fmt.Scanln(&date)
	fmt.Scanln(&first_name)
	fmt.Scanln(&second_name)
	fmt.Scanln(&surname)
	fmt.Scanln(&cost_one)
	fmt.Scanln(&cost_two)
	fmt.Scanln(&cost_third)

	t, _ := time.Parse("02.01.2006", date)

	totalSum := cost_one + cost_two + cost_third
	totalKopecks := math.Round(totalSum * 100)
	rub := int(totalKopecks) / 100
	kop := int(totalKopecks) % 100

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %v. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп.\n\nС уважением,\nГл. бух. Иванов А.Е.\n",
		second_name,
		first_name,
		surname,
		t.Add(time.Hour*24*15).Format("02.01.2006"),
		rub,
		kop)
}
