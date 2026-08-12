package main

import "fmt"

func main() {
    var thing string
    fmt.Scanln(&thing)
    fmt.Println(fmt.Sprintf("А ещё я люблю %s!", thing))
}