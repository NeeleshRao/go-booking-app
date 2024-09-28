package main

import "fmt"

func main() {

	var confName string = "Go Conference"
	const confTickets int = 50
	var remainingTickets uint32 = 50
	var booking []string

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

	remainingTickets -= userTickets

	booking = append(booking, firstName+" "+lastName)

	fmt.Printf("The booking slice is %v\n", booking)
	fmt.Printf("The first element of the slice is %v\n", booking[0])
	fmt.Printf("The type of the slice is %T\n", booking)
	fmt.Printf("The length of the slice is %v\n", len(booking))

	fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("A confirmation mail will be sent to your email - %v\n", email)
	fmt.Printf("For the event %v, only %v seats remain!\n", confName, remainingTickets)
}
