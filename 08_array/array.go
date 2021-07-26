package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad Yusron"
	names[1] = "Hartoyo"
	names[2] = "Yusron"

	var firstName = names[0]
	var lastName = names[1]
	var nickname = names[2]

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(nickname)
	// Jumlah array tidak bisa bertambah
	var nilai = [3]int{
		87, 98, 89,
	}
	// Mengubah nilai array index ke 2
	nilai[2] = 100
	fmt.Println(nilai)
	// Menghitung panjang array
	fmt.Println(len(nilai))
}
