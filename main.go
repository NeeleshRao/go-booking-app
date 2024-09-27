package main

import "fmt"

func main() {
	// var confName = "Python Convention"
	// const confTickets = 50
	// var remainingTickets = 50
	// another operator for type inferring := but this cannot be done outside function
	// other ways of declaration with types rather than inferring
	var confName string = "Go Conference"
	const confTickets int = 50
	var remainingTickets uint32 = 50

	// To view types of the variables
	fmt.Printf("The types of variables confName = %T, confTickets = %T and remainingTickets = %T\n", confName, confTickets, remainingTickets)

	fmt.Printf("Welcome to the %v conference!\n", confName)
	fmt.Printf("There are a total of %v tickets with only %v remaining...\n", confTickets, remainingTickets)

	var userName string
	var numTickets uint

	userName = "neelesh ig"
	numTickets = 2

	fmt.Printf("User %v wants %v tickets.\n", userName, numTickets)
}
