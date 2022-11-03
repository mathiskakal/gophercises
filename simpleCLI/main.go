package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Defining Package level Variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// this creates a list of maps that have strings as key and value
// the 0 is the initial size of the list
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// part 1/4 of synchro : create a waitGroup to make the program wait for the goroutine to end to exit
var wg = sync.WaitGroup{}

// Main Function
func main() {

	// Welcome Messages
	greetUsers()

	// loop logic
	// Get user input for their first name, last name, email and validate their input
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	// if all conditions met, then execute the booking
	if isValidName && isValidEmail && isValidTicketNumber {
		// Adding info to array and counting remaining tickets
		// remainingTickets, bookings = bookTicket(userTickets, firstName, lastName, email)
		bookTicket(userTickets, firstName, lastName, email)

		// part 2/4 of synchro :
		// Add that 1 goroutine to the wait group
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		// Print the first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookigns are: %v\n", firstNames)

		// Logic for exiting app
		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year.\n")

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
	// part 3/4 of synchro : tell the main logic to actually wait for queued elements
	wg.Wait()
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// make sure that input name and last name contain at least two characters. Will be true if condition is met
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// same thing for email validation, we make sure it contains @, will return true if it does
	isValidEmail := strings.Contains(email, "@")

	// validate that both user requested a positive whole number of tickets and that it doesn't exceed the amount left.
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############")
	// part 4/4 of synchro : be able to detect when we are actually done waiting
	wg.Done()

}
