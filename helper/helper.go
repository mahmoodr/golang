package main

import "strings"

func inputValidation(Name string, Family string, email string, ticket int, remaining int) (bool, bool, bool) {
	isValidName := len(Name) >= 2 && len(Family) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticket > 0 && ticket <= remaining
	return isValidName, isValidEmail, isValidTicketNumber
}
