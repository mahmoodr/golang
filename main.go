package main

import (
	"fmt"
)

func main() {
	var conferanceName = "GoLang conferance"
	const conferanceTicket = 50
	var remainingTicket int = 50

	fmt.Printf("\n\n\nWelcome to %v booking application system.\n", conferanceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferanceTicket, remainingTicket)
	fmt.Printf("Please first book your ticket to participate in the conferance.\n\n\n")

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

	remainingTicket = remainingTicket - userTicket
	fmt.Printf("the remaining tickets is: %v\n", remainingTicket)

	fmt.Printf("welcome %v %v and thanks for purchasing %v tickets. you'll recieve confirmation email soon!\n", userName, userFamily, userTicket)

	//fmt.Printf("Data Type of conferanceName is: %T \n conferanceTicket is: %T \n remainingTicket is: %T\n userName is: %T \n userFamily is: %T", conferanceName, conferanceTicket, remainingTicket, userName, &userFamily)
	//fmt.Println(userFamily)

	//Arrays and slices
	var booking = []string{} // define a slice
	booking = append(booking, userName+" "+userFamily)
	fmt.Println(booking)

}
