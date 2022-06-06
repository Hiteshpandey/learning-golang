// Created a new package helper
package helper

import "strings"

// First letter of the function capitalized so as to be exportable to other packages
func ValidateUserInput(firstName string, lastName string, email string, UserTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValideEmail := strings.Contains(email, "@") && len(email) > 12
	isValidTicketNumber := UserTickets > 0 && UserTickets <= remainingTickets
	// isValidCity := city == "Singapore" || city == "London"
	return isValidName, isValideEmail, isValidTicketNumber
}
