package main

import "fmt"

func main() {
	var confName = "Python Convention"
	const confTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v conference!\n", confName)
	fmt.Printf("There are a total of %v tickets with only %v remaining...", confTickets, remainingTickets)
}
