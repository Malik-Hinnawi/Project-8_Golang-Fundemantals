package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createrTransformer(2)
	transformed := transformNumbers(&numbers, double)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := make([]int, len(*numbers))

	for i, v := range *numbers {
		dNumbers[i] = transform(v)
	}

	return dNumbers
}

// Factory Design Pattern
func createrTransformer(factor int) func(int) int {
	// Factor is a closure parameter
	return func(n int) int {
		return n * factor
	}
}
