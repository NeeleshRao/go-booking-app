package main

import "strings"

func validateInput(firstName string, lastName string, email string, userTickets uint32, remainingTickets uint32) (bool, bool, bool) {
	// validating user input
	// to run multiple files use git run .
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTickets = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
