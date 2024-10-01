package main

import (
	"fmt"
	"strings"
)

var confName string = "Go Conference"

const confTickets int = 50

var remainingTickets uint32 = 50
var booking []string

func main() {

	greeting(confName, confTickets, remainingTickets)

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := userInput()

		isValidName, isValidEmail, isValidTickets := ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTickets(firstName, lastName, email, userTickets)
			dispAllBookings(booking)

		} else {
			displayError(isValidName, isValidEmail, isValidTickets)
		}

	}

}

func greeting(confName string, confTickets int, remainingTickets uint32) {
	fmt.Printf("Welcome to the %v conference!\n", confName)
	fmt.Printf("There are a total of %v tickets with only %v remaining...\n", confTickets, remainingTickets)
}

func userInput() (string, string, string, uint32) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint32

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email id : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func displayError(isValidName bool, isValidEmail bool, isValidTickets bool) {
	if !isValidName {
		fmt.Println("You have entered a very short name (min 2 characters)!")
	}
	if !isValidEmail {
		fmt.Println("You did not enter a valid email!")
	}
	if !isValidTickets {
		fmt.Printf("There are only %v tickets remaining, please try again!\n", remainingTickets)
	}
}

func bookTickets(firstName string, lastName string, email string, userTickets uint32) {
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation mail will be sent to your email - %v\n", email)
	fmt.Printf("For the event %v, only %v seats remain!\n", confName, remainingTickets)

	booking = append(booking, firstName+" "+lastName)
}

func ValidateInput(firstName string, lastName string, email string, userTickets uint32, remainingTickets uint32) (bool, bool, bool) {

	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTickets = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func dispAllBookings(booking []string) {
	var onlyFirstNames []string

	for _, book := range booking {
		onlyFirstNames = append(onlyFirstNames, strings.Split(book, " ")[0])
	}

	fmt.Printf("These are all the bookings till date : %v\n", onlyFirstNames)
}
