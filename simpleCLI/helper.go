package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// make sure that input name and last name contain at least two characters. Will be true if condition is met
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// same thing for email validation, we make sure it contains @, will return true if it does
	isValidEmail := strings.Contains(email, "@")

	// validate that both user requested a positive whole number of tickets and that it doesn't exceed the amount left.
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
