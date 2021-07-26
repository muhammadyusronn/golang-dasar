package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // dia akan error karena nilai maks adalah 127 sedangkan 100000 sudah melebihi

	var name = "Yusron"
	var e byte = name[0]
	// conversi byte ke string
	var eString string = string(e)
	fmt.Println("Ini adalah byte ", e)
	fmt.Println("ini mengknversi byte ke string ", eString)
}
