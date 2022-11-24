package main

import "fmt"

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var reamainingTickets uint = conferenceTickets
	var bookings [50]string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, reamainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, reamainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	reamainingTickets = reamainingTickets - userTickets

	fmt.Printf("The whole array ")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", reamainingTickets, conferenceName)
}
