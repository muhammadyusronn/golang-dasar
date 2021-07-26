package main

import "fmt"

func main() {
	// Membuat tipe data baru dari tipe data yang sudah ada
	type noKTP string
	var yusronKTP noKTP = "09010581620013"
	fmt.Println(yusronKTP)
}
