package main

import "fmt"

func main() {
	var productNames [4]string
	productNames[0] = "A Book"

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames[0])

	featuredPrices := prices[1:3]
	featuredPrices[0] = 199.9
	highlightedPrices := featuredPrices[:1]

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = featuredPrices[:3]

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
