package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Dog struct{}
type Cat struct{}

func (_ Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (_ Cat) MakeSound() {
	fmt.Println("Мяу!")
}
