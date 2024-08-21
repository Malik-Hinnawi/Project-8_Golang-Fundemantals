package main

import "fmt"

func main() {
	prices := []float64{10.99, 8, 99}
	prices = append(prices, 5.9, 6.999)
	fmt.Println(prices)

	discountPrices := []float64{101, 99, 80.99, 20.59}

	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
