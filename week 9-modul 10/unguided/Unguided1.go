package main

import "fmt"

type arrFloat [1000]float64

func main() {
	var A arrFloat
	var n int
	var min, max float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	min = A[0]
	max = A[0]

	for i := 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}

	fmt.Printf("%.2f %.2f\n", min, max)
}