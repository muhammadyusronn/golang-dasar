package main

import "fmt"

func main() {
	// Multiple Declaration
	const (
		firstName string = "Muhammad Yusron"
		lastName         = "Hartoyo"
	)
	const value = 1000
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	// Constant tidak bisa di rubah nilainya, beda seperti variabel
	// firstName = "hehehe" (Akan error)
}
