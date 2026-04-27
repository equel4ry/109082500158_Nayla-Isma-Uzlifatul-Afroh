package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func main() {
	var data [NMAX]int
	var n int = 7

	// isi array
	data[0] = 10
	data[1] = 15
	data[2] = 20
	data[3] = 25
	data[4] = 30
	data[5] = 35
	data[6] = 40

	// a. tampil semua
	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// b. indeks ganjil
	fmt.Print("Indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	// c. indeks genap
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	// d. kelipatan x
	var x int = 2
	fmt.Print("Kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	// e. hapus indeks 2
	idx := 2
	for i := idx; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n--

	fmt.Print("Setelah hapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// f. rata-rata
	total := 0
	for i := 0; i < n; i++ {
		total += data[i]
	}
	rata := float64(total) / float64(n)
	fmt.Printf("Rata-rata: %.2f\n", rata)

	// g. standar deviasi
	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rata
		jumlah += selisih * selisih
	}
	std := math.Sqrt(jumlah / float64(n))
	fmt.Printf("Standar deviasi: %.2f\n", std)

	// h. frekuensi
	cari := 25
	freq := 0
	for i := 0; i < n; i++ {
		if data[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi", cari, ":", freq)
}
