package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := performCalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	writeToFile(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		return -1, errors.New("input should not be less than 0")
	}

	if userInput == 0 {
		return -1, errors.New("input should not be 0")
	}

	return userInput, nil
}

func performCalculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeToFile(ebt, profit, ratio float64) {
	resultsText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}
