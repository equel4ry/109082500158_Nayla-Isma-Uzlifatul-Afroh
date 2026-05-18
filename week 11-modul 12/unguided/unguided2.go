package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	masuk := 0
	sah := 0

	fmt.Scan(&x)
	for x != 0 {
		masuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}

		fmt.Scan(&x)
	}

	ketua := -1
	wakil := -1

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			if ketua == -1 || suara[i] > suara[ketua] || (suara[i] == suara[ketua] && i < ketua) {
				wakil = ketua
				ketua = i
			} else if wakil == -1 || suara[i] > suara[wakil] || (suara[i] == suara[wakil] && i < wakil) {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}