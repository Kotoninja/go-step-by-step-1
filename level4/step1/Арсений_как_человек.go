package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) PrettyPrint() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s", p.name, p.age, p.address)
}

func main() {
	man := Person{name: "Anton", age: 31, address: "Krasnogorsk"}
	man.PrettyPrint()
}