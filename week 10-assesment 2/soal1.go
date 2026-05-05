package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	found := false
	i := 0
	for i < n && !found {
		if T[i] == val {
			found = true
		}
		i++
	}
	return found
}

func inputSet(T *set, n *int) {
	var val int
	*n = 0
	fmt.Scan(&val)
	for !exist(*T, *n, val) && *n < 2022 {
		(*T)[*n] = val
		*n++
		fmt.Scan(&val)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			(*T3)[*h] = T1[i]
			*h++
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Print(T[i], " ")
		} else {
			fmt.Print(T[i])
		}
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}