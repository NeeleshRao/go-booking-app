package helper

import "strings"

func ValidateInput(firstName string, lastName string, email string, userTickets uint32, remainingTickets uint32) (bool, bool, bool) {

	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTickets = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
