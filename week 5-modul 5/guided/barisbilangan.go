package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	baris(n)
}

func baris(n int) {
	if n == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(n)
		baris(n - 1)
	}
}
