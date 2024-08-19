package main

import "fmt"

func main() {
	prices := []float64{10.99, 8, 99}
	prices = append(prices, 5.99)
	fmt.Println(prices)

}
