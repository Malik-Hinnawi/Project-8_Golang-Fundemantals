package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("Can't continue sorry")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid Format")
			return
		}

		fmt.Printf("Your choice: %v\n", choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		case 2:
			fmt.Print("Enter deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated! new amount: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Enter withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Must be less than the account balance.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Balance updated! new amount: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}

}
