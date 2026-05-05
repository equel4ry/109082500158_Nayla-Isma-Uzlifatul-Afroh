package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&(*T)[i].NIM, &(*T)[i].nama, &(*T)[i].nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	i := 0
	for i < n && T[i].NIM != nim {
		i++
	}
	if i < n {
		return T[i].nilai
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	maks := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == nim && T[i].nilai > maks {
			maks = T[i].nilai
		}
	}
	return maks
}

func main() {
	var T arrayMahasiswa
	var n int
	var nim string

	inputData(&T, &n)
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nim)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nim, cariNilaiPertama(T, n, nim))
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nim, cariNilaiTerbesar(T, n, nim))
}