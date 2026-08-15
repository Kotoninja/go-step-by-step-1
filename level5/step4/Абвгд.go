package main

import (
	"fmt"
	"slices"
)

func SortNames(names []string) {
	slices.Sort(names)
}

func main() {
	arr := []string{"Аксинья", "Варвара", "Арина", "Есения"}
	SortNames(arr)

	fmt.Println(arr)
}
