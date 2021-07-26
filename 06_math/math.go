package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var result = a + b
	// atau
	var result2 = 0
	result2 += b

	fmt.Println(result)
	fmt.Println(result2)
}
