package main

import "fmt"

func main() {
	f1 := factorial(5)
	fmt.Println(f1)

	f2 := factorialIterative(5)
	fmt.Println(f2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialIterative(n int) int {
	result := 1

	for i := 0; i < n; i++ {
		result *= i
	}
	return result
}
