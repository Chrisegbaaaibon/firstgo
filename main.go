package main

import "fmt"

func main () {
	conferenceName  := "Golang Conference"
	fmt.Printf("Welcome to %s\n", conferenceName)
	
	var firstName string
	var lastName string
	var email string
	var age int
	var userTickets int
	var totalTickets int = 50
	var remainingTickets int

	//Excecution
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Printf("You are not eligible to attend the conference\n")
	} else {
		 
	}

}