package main

import "fmt"

type transformFN func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{2, 3, 4, 5}

	tf1 := getTransformerFunction(&numbers)
	tf2 := getTransformerFunction(&moreNumbers)
	numberstf1 := transformNumbers(&numbers, tf1)
	numberstf2 := transformNumbers(&moreNumbers, tf2)

	fmt.Println(numbers)
	fmt.Println(numberstf1)
	fmt.Println(numberstf2)
}

func getTransformerFunction(numbers *[]int) transformFN {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func transformNumbers(numbers *[]int, transform transformFN) []int {
	dNumbers := make([]int, len(*numbers))

	for i, v := range *numbers {
		dNumbers[i] = transform(v)
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
