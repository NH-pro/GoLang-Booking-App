package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50 // int type
	var remainingTickets uint = 50   // uint is another int type
	// var bookings [50]string          // Declaring an array of strings that can hold up to 50 indexes
	var bookings []string // Declaring a "slice", an array without a index limit constraint

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets:")
		fmt.Scan(&userTickets)

		/**
		The "&" is used to reference the "pointer" for the "userName" var.
		Pointers are the memory address where the data is stored.
		So here, the app stops and the user is allowed to enter in data for the var.
		*/

		if userTickets > remainingTickets {
			fmt.Printf("We onyl have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our confernece is booked out. Come back next year. ")
			break
		}
	}
}
