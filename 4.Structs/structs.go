package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Struct literal decleration:
	appUser, err := user.New(firstName, lastName, birthdate)
	// ... do something awesome with that gathered data!

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@delta.smart", "123456")
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
