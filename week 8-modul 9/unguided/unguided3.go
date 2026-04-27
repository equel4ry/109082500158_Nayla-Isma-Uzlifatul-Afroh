package main

import "fmt"

const NMAX = 100

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [NMAX]string
	var n int = 0
	var i int = 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Print("Pertandingan ", i, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}

		n++
		i++
	}

	for j := 0; j < n; j++ {
		fmt.Println("Hasil", j+1, ":", hasil[j])
	}

	fmt.Println("Pertandingan selesai")
}