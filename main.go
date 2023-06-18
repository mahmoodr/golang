package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferanceName = "GoLang conferance"
	const conferanceTicket = 50
	var remainingTicket int = 50

	greetuser(conferanceName)

	fmt.Printf("\n\n\nWelcome to %v booking application system.\n", conferanceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferanceTicket, remainingTicket)
	fmt.Printf("Please first book your ticket to participate in the conferance.\n\n\n")

	for {

		var userName string
		var userFamily string
		var userEmail string
		var userPhone string
		var userTicket int

		//ask for user input
		fmt.Printf("Please Enter your personal data to register.\n")

		fmt.Printf("First name:\n")
		fmt.Scan(&userName)

		fmt.Printf("Family name:\n")
		fmt.Scan(&userFamily)

		fmt.Printf("Email:\n")
		fmt.Scan(&userEmail)

		fmt.Printf("Phone:\n")
		fmt.Scan(&userPhone)

		fmt.Printf("Number of ticket:\n")
		fmt.Scan(&userTicket)

		inputValidation(userName, userFamily, userEmail, userTicket, remainingTicket)

		remainingTicket = remainingTicket - userTicket

		var userData = make(map[string]string)

		if inputValidation() {
			fmt.Printf("the remaining tickets is: %v\n", remainingTicket)

			fmt.Printf("welcome %v %v and thanks for purchasing %v tickets. you'll recieve confirmation email soon!\n", userName, userFamily, userTicket)

			//fmt.Printf("Data Type of conferanceName is: %T \n conferanceTicket is: %T \n remainingTicket is: %T\n userName is: %T \n userFamily is: %T", conferanceName, conferanceTicket, remainingTicket, userName, &userFamily)
			//fmt.Println(userFamily)

			//Arrays and slices//
			var booking = []string{} // define a slice//
			booking = append(booking, userName+" "+userFamily)
			fmt.Println(booking)

			firstNames := []string{}
			for _, booking := range booking {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("the first names %v\n", firstNames)

			// loop to check the remaining tickets and stop when it's finished
			if remainingTicket == 0 {
				//end program
				fmt.Println("our conferance is sold out!")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Printf("Email not valid!")
			}
			if !isValidName {
				fmt.Printf("first or last name not valid")
			}
			if !isValidTicketNumber {
				fmt.Printf("requested ticket number not valid")
			}
		}

	}

	city := "london"
	switch city {
	case "new york", "istanbul":
		//new yourk booking
	case "singapure", "hong kong":
		//singapore booking
	default:
		fmt.Printf("no valid city")
	}

}

func greetuser(conferanceName string) {
	fmt.Println("welcome to confrance!")
}

func inputValidation(Name string, Family string, email string, ticket int, remaining int) (bool, bool, bool) {
	isValidName := len(Name) >= 2 && len(Family) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticket > 0 && ticket <= remaining
	return isValidName, isValidEmail, isValidTicketNumber
}
