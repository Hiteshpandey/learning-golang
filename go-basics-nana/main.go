package main

import (
	"fmt"
	"go-basics-nana/helper"
	"sync"
	"time"
)

// Package level varaiables, available globally in the package
var conferenceName string = "Go conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets

//[] before map means it creates a list of maps instead of a single map
// [ {k1:v1, k2,v2},{k4:v4, k5:v5} ]
// 0 passed means the initial size of the list that would be zero elements
// var bookings = make([]map[string]string, 0)

// list of UserData struct type
//[] before UserData means it creates a list of structs instead of a single struct
// 0 passed means the initial size of the list, that would be zero elements
// this creates an empty slice list of type UserData struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// var city string

	// Predefined set of values, size of array
	// var bookings [50]string

	greetUsers()

	// for remainingTickets > 0 {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValideEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidTicketNumber && isValideEmail {

		bookTicket(userTickets, firstName, lastName, email)
		// We are adding one thread to wait group
		// This is basically adding a counter of how many number of thread(s) the main thread should wait for before exiting
		wg.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("All of the bookings till now: %v.\n", getFirstNames())

		if remainingTickets == 0 {
			fmt.Println("All tickets are booked!!!. If you want new reservations try at the next event.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("Please enter valid first or last name")
		}

		if !isValideEmail {
			fmt.Println("Please enter valid email")
		}

		// if !isValidCity {
		// 	fmt.Println("Please enter valid city")
		// }

		if !isValidTicketNumber {
			fmt.Println("We don't have booking for the provided number of tickets")
		}
	}

	// Wait for all threads to end afther reaching this point/line of code
	wg.Wait()
}

// }

func greetUsers() {
	fmt.Printf("\nWelcome to our %v booking system.\n", conferenceTickets)
	fmt.Printf("We have total of %v tickets, and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets booked to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	// range can be used to iterate over different data structures
	// iterate over the array and get only first name by splitting string
	for _, booking := range bookings {
		// values can be defined with assignment without the type only when being assigned.
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their name
	// Need to take value by reference in scan since the value needs to be mutated
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you require: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for the user details
	// In go map can only have same datatype for key value
	// Key value types cannot be different
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// Convert int to unit 64 bit type, then convert to string, with base value 10. Which means decimal value.
	// userData["number"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v ticktes remains for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Sprintf returns a formatted string instead of printing it
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("################################")

	// Remove the thread counter from waiting list
	wg.Done()
}
