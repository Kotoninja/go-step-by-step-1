package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, z, m, n float64

	fmt.Scan(&x, &y, &z, &m, &n)

	// numerator := 5.0 * x * math.Sin(2*y)
	// denominator := z + math.Pow(n, math.Log(m))
	// result := numerator / denominator

	fmt.Println((5.0 * x * math.Sin(2*y)) / (z + math.Pow(n, math.Log(m))))
}
