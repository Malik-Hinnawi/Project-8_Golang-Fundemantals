package main

import "fmt"

// Type alias:
type floatMap map[string]float64

func (fm floatMap) output() {
	fmt.Println(fm)
}

func main() {
	// type, size, cap
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Test")
	userNames = append(userNames, "Test 1")
	userNames = append(userNames, "Test 2")

	// Now copy array:
	userNames = append(userNames, "Test 3")
	userNames = append(userNames, "Test 4")

	fmt.Println(userNames)

	// type, cap only!
	// Using type alias
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	// Now copy map:
	courseRatings["godot"] = 4.6

	courseRatings.output()

	for i, v := range userNames {
		fmt.Println(i, ":", v)
	}

	for k, v := range courseRatings {
		fmt.Println(k, "rating:", v)
	}
}
