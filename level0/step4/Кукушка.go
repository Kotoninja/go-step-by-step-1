package main

import (
	"fmt"
	"time"
)

func main() {
	var input_time string

	fmt.Scanln(&input_time)

	t, _ := time.Parse("2006-01-02/15:04:05", input_time)

	s := fmt.Sprintf("Текущее время %v часов, %v минут. Ты точно не забыл про важные дела на сегодня?", t.Hour(), t.Minute())

	fmt.Println(s)
}
