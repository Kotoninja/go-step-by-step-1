package main

import (
	"math"
)

func findDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c

}

func SquareRoots(a, b, c float64) (float64, float64) {
	discriminant := findDiscriminant(a, b, c)

	// fmt.Println(discriminant)
	if discriminant < 0 {
		return 0, 0
	}

	var first float64 = (-b + math.Sqrt(discriminant)) / (2 * a)
	var second float64 = (-b - math.Sqrt(discriminant)) / (2 * a)

	return math.Min(first, second), math.Max(first, second)
}
