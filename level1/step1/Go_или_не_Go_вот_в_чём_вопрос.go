package main

import (
    "fmt"
)

func main() {
    var name string

    fmt.Scanln(&name)

    if name == "Go" {
    	fmt.Println("Go!")
        return
    }
    fmt.Println("Я знаю только Go!")
}