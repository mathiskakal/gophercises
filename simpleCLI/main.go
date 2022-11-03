package main

import (
	"fmt"
	"strings"
)

// Defining Package level Variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

// Main Function
func main() {

	// Welcome Messages
	greetUsers()

	// loop logic
	for {
		// Get user input for their first name, last name, email and validate their input
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// if all conditions met, then execute the booking
		if isValidName && isValidEmail && isValidTicketNumber {
			// Adding info to array and counting remaining tickets
			remainingTickets, bookings = bookTicket(userTickets, firstName, lastName, email)

			// Print the first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookigns are: %v\n", firstNames)

			// Logic for exiting app
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			// Here we chain if statements to check if any conditions were met to assist the user on the validation of their input
			if !isValidName {
				fmt.Println("The firstname or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain an '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}
		}
	}
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want ?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// iterating in bookings list to display only first names for more privacy
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
