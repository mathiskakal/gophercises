package main

import (
	"fmt"
	"strings"
)

func main() {
	//  defining constants
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// Welcome message
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// loop logic
	for {
		// Declaring Variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their first name, last name, email
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want ?")
		fmt.Scan(&userTickets)

		// make sure that input name and last name contain at least two characters. Will be true if condition is met
		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		// same thing for email validation, we make sure it contains @, will return true if it does
		isValidEmail := strings.Contains(email, "@")

		// validate that both user requested a positive whole number of tickets and that it doesn't exceed the amount left.
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// if all conditions met, then execute the booking
		if isValidName && isValidEmail && isValidTicketNumber {
			// Adding info to array and counting remaining tickets
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// iterating in bookings list to display only first names for more privacy
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Logic for exiting app
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			/*
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
				// will restart the loop skipping what is afer this if statement
			*/
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
