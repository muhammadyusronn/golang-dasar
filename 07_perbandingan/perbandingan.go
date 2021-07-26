package main

import "fmt"

func main() {
	var firstname = "Muhammad Yusron"
	var lastname = "Hartoyo"

	var result bool = firstname == lastname
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
