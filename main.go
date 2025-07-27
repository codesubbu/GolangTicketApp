package main

import "fmt"

func main() {
	var FirstName string
	var LastName string
	var Email string
	var RemSlots int = 50

	var Slots int

	fmt.Println("Enter your first name ")
	fmt.Scan(&FirstName)

	fmt.Println("Enter your last name ")
	fmt.Scan(&LastName)

	fmt.Println("Enter your Email ")
	fmt.Scan(&Email)

	fmt.Println("Enter Slots ")
	fmt.Scan(&Slots)

	RemSlots = RemSlots - Slots
	fmt.Printf("Thanks you %v %v for booking %v slots and confirmation has sent to %v\n", FirstName, LastName, Slots, Email)
	fmt.Println("The remaining slots are ", RemSlots)

}
