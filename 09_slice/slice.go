package main

import "fmt"

func main() {
	// Slice merupakan tipe data potongan dari array bedanya adalah slice ukurannya bisa berubah sedangkan array tidak bisa
	// Function slice
	// len(slice) mendapatkan panjang slice
	// cap(slice) mendapatkan kapasitas slice
	// append(slice, data) membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh akan membuat array baru
	// make([] TypeData, length, capacity) membuat slice baru
	// copy(destionation, source) menyaling slice dari source ke destination
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:7]
	// Menubah slice
	slice1[0] = "Ubah"
	// MENCETAK slice
	fmt.Println(slice1)
	// Panjang slice
	fmt.Println(len(slice1))
	// Kapasitas slice
	fmt.Println(cap(slice1))

	// Menambahkan data slice
	var slice2 = bulan[10:]
	// Mencetak slice 2
	fmt.Println(slice2)
	var slice3 = append(slice2, "Bulan Baru")
	// Mencetak slice yang sudah ditambah data
	fmt.Println(slice3)
	// Mengubah nilai slice3 index ke 1
	slice3[1] = "Bukan Desember"
	// Mencetak hasil slice 3
	fmt.Println(slice3)
	// Mencetak slice 2
	fmt.Println(slice2)
	// Mencetak array
	fmt.Println(bulan)
	// Data array dan slice 2 tidak berubah karen sudah melebihi kapasitas sehingga membuat array baru
	// Membuat slice baru
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Muhammad Yusron"
	newSlice[1] = "Hartoyo"
	fmt.Println(newSlice)
	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	// function copy slice
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	// Perberdaan declarasi slice dan array
	iniArray := [3]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	fmt.Println("ini adalah array -> ", iniArray)
	fmt.Println("ini adalah slice -> ", iniSlice)

}
