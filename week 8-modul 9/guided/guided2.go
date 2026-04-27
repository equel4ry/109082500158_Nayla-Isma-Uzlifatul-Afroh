package main

import "fmt"

func main() {
	// deklarasi map
	nilai := make(map[string]int)

	// tambah data
	nilai["Dhimas"] = 90
	nilai["Ichya"] = 90

	// tampilkan isi map
	fmt.Println("Data nilai:")
	for nama, grade := range nilai {
		fmt.Println(nama, "=", grade)
	}

	// update nilai
	nilai["Ichya"] = 80

	// searching
	cariNama := "hafizh"
	isiData, ok := nilai[cariNama]

	if ok {
		fmt.Println("Nilai", cariNama, "=", isiData)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	// hapus data
	delete(nilai, "Dhimas")

	// tampilkan lagi setelah perubahan
	fmt.Println("\nData setelah update & delete:")
	for nama, grade := range nilai {
		fmt.Println(nama, "=", grade)
	}
}