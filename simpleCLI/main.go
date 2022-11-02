package main

import (
	"fmt"
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

		// Adding info to array and counting remaining tickets
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}
