package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Normal version:
	sum := sumup(numbers)
	fmt.Println("Slice sum result:", sum)

	// Variadic version:
	sum = sumupVariadic(1, 2, 3)
	fmt.Println("Variadic result:", sum)

	// Passing slices into variadic functions:
	sum = sumupVariadic(numbers...)
	fmt.Println("Passing slice result:", sum)
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumupVariadic(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
