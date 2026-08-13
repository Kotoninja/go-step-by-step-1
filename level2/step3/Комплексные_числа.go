package main

import (
	"fmt"
	// "math/cmplx"
)

func PrintComplexNumber(z complex64) {
	re := real(z)
	im := imag(z)

	fmt.Println(fmt.Sprintf("Действительная часть: %.2f. Мнимая часть: %.2f", re, im))
}

func main() {
	PrintComplexNumber(complex(10, 15))

}
