package main

import "fmt"

type arrFloat [1000]float64

func main() {
	var A arrFloat
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&A[i])
	}

	var jumlahWadah int
	if x%y == 0 {
		jumlahWadah = x / y
	} else {
		jumlahWadah = (x / y) + 1
	}

	var hasil [1000]float64
	idx := 0

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0
		for j := 0; j < y && idx < x; j++ {
			total += A[idx]
			idx++
		}
		hasil[i] = total
	}

	// output baris 1
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", hasil[i])
	}
	fmt.Println()

	// rata-rata
	sum := 0.0
	for i := 0; i < jumlahWadah; i++ {
		sum += hasil[i]
	}

	rata := sum / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rata)
}