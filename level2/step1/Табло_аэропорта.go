package main

import "fmt"

func PrintFlightRow(number_a_flight string, city_from, city_to string, flight_duration float64, reception_desk, gate int, is_registration bool) {
    if !is_registration {
        fmt.Printf("| %s | %s—%s | %d регистрация продолжается |\n", number_a_flight, city_from, city_to, reception_desk)
        return
    }
    fmt.Printf("| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %0.1f часа |\n", number_a_flight, city_from, city_to, gate, flight_duration)
}