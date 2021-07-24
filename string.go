package main

import "fmt"

func main() {
	var fullname string
	fullname = "Muhammad Yusron Hartoyo"
	nickname := "Yusron"

	fmt.Println(nickname)
	fmt.Println(fullname)
	// Menghitung jumlah string
	fmt.Println(len(nickname))
	// Mengambil karakter berdasarkan posisi yang ditentukan
	fmt.Println(fullname[1])
}
