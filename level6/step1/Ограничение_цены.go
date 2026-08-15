package main

const (
	minPrice float64 = 99
	maxPrice float64 = 20000
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	} else if price > maxPrice {
		return maxPrice
	} else {
		return price
	}
}
