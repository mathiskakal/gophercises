package main

import (
	"fmt"
)

func main() {
	//  defining constants
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// Welcome message
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Declaring Variables
	var bookings [50]string

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
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The type of array: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
