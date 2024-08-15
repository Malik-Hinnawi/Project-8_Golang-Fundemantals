package main

import "fmt"

func main() {
	age := 32

	// Referencing
	agePointer := &age

	// Dereferncing:
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(&age)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
