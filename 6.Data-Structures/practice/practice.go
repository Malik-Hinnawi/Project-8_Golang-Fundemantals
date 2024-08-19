package main

import (
	"fmt"

	"examples.com/practice/product"
)

func main() {
	// 1)
	hobbies := [3]string{"Badminton", "Swimming", "Table Tennis"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println("First element:", hobbies[0])
	fmt.Println("Second/Third element:", hobbies[1:])

	// 3)
	slice1 := hobbies[:2]
	slice2 := hobbies[0:2]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 4)
	slice2 = hobbies[1:3]
	fmt.Println(slice2)

	// 5)
	goals := []string{"Goal 1", "Goal 2"}

	// 6)
	goals[1] = "Goal 2 updated"
	goals = append(goals, "Goal 3")
	fmt.Println(goals)

	// 7)

	product1 := product.New("Product 1", 1, 6.44)
	product2 := product.New("Product 2", 2, 2.99)
	products := []product.Product{product1, product2}

	product3 := product.New("Product 3", 3, 3.99)
	products = append(products, product3)
	products[0].Display()

}
