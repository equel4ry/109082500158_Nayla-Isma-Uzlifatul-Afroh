package main
import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	// volume = π × r² × t
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	// massa = volume × massa jenis
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 > m2 {
		fmt.Println("Selisih massa zat cair kiri dan massa zat cair kanan :", m1-m2)
	} else {
		fmt.Println("Selisih massa zat cair kiri dan massa zat cair kanan :", m2-m1)
	}
}

func main() {
	var r float64 // jari-jari
	var tKiri, tKanan float64 // tinggi
	var mjKiri, mjKanan float64 // massa jenis
	var massaKiri, massaKanan float64 // massa

	// masukkan jari-jari
	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	// input kiri
	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	// input kanan
	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	// hitung massa
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	// tampilkan hasil
	display(massaKiri, massaKanan)
}