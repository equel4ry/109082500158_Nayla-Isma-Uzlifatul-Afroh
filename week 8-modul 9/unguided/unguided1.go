package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	px, py int
	r      int
}

func hitungJarak(a, b titik) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func cekDalam(l lingkaran, p titik) bool {
	jarak := hitungJarak(titik{l.px, l.py}, p)
	return jarak <= float64(l.r)
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Scan(&l1.px, &l1.py, &l1.r)
	fmt.Scan(&l2.px, &l2.py, &l2.r)
	fmt.Scan(&p.x, &p.y)

	c1 := cekDalam(l1, p)
	c2 := cekDalam(l2, p)

	if c1 && c2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if c1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if c2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}