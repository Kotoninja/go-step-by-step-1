package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func TimeNow() time.Time {
	return time.Date(2026, time.August, 13, 11, 30, 30, 0, time.UTC)
}

func fromRuToEng(phrase string) string {
	switch strings.ToLower(phrase) {
	case "понедельник":
		return "mondey"
	case "вторник":
		return "tuesday"
	case "среда":
		return "wednesday"
	case "четверг":
		return "thursday"
	case "пятница":
		return "friday"
	case "суббота":
		return "saturday"
	case "воскресенье":
		return "sunday"
	default:
		return ""
	}
}

func fromEngToRu(phrase string) string {
	switch strings.ToLower(phrase) {
	case "mondey":
		return "понедельник"
	case "tuesday":
		return "вторник"
	case "wednesday":
		return "среда"
	case "thursday":
		return "четверг"
	case "friday":
		return "пятница"
	case "saturday":
		return "суббота"
	case "sunday":
		return "воскресенье"
	default:
		return ""
	}
}
func dayOfWeek() time.Weekday {
	return TimeNow().Weekday()
}

func currentDayOfTheWeek() string {
	return strings.Title(fromEngToRu(dayOfWeek().String()))
}

func dayOrNight() string {
	currentTime := TimeNow()
	if currentTime.Hour() >= 10 && currentTime.Hour() <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	difference := 5 - int(dayOfWeek())
	if difference >= 0 {
		return difference
	}
	return 7 + difference

}

func CheckCurrentDayOfTheWeek(answer string) bool {
	return fromRuToEng(strings.ToLower(answer)) == strings.ToLower(dayOfWeek().String())
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	return strings.EqualFold(strings.ToLower(answer), strings.ToLower(dayOrNight())), nil
}

func main() {
	// fmt.Println(currentDayOfTheWeek())
	// fmt.Println(dayOrNight())
	// nextFriday()
	// fmt.Println(nextFriday())
	// fmt.Println(CheckCurrentDayOfTheWeek("Четверг"))
	// fmt.Println(CheckCurrentDayOfTheWeek("Четверг"))
	fmt.Println(CheckNowDayOrNight("день"))
	fmt.Println(CheckNowDayOrNight("День"))
	fmt.Println(CheckNowDayOrNight("ДеНь"))
	fmt.Println(CheckNowDayOrNight("ночь"))
	fmt.Println(CheckNowDayOrNight("ок"))
	fmt.Println(CheckNowDayOrNight("гидроэлектростанция"))
}
