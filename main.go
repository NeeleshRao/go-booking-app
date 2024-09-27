package main

import "fmt"

func main() {
	
	var confName string = "Go Conference"
	const confTickets int = 50
	var remainingTickets uint32 = 50

	// to understand pointers in go, it is very simple. 
	// Every variable basically references where a value is stored.
	// A pointer holds the address of where the value is/will be stored in memory
	fmt.Println(confName, &confName)
	
	fmt.Printf("Welcome to the %v conference!\n", confName)
	fmt.Printf("There are a total of %v tickets with only %v remaining...\n", confTickets, remainingTickets)

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

	fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation mail will be sent to your email - %v\n", email)
}
