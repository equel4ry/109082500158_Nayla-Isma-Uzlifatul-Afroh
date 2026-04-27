package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	var ch rune

	fmt.Scanf("%c", &ch)
	for ch != '.' && *n < NMAX {
		t[*n] = ch
		*n++
		fmt.Scanf("%c", &ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// isi array
	isiArray(&tab, &n)

	// cetak awal
	cetakArray(tab, n)

	// cek palindrom
	if palindrom(tab, n) {
		fmt.Println("palindrom")
	} else {
		fmt.Println("bukan palindrom")
	}

	// balik array
	balikanArray(&tab, n)

	// cetak hasil balik
	cetakArray(tab, n)
}