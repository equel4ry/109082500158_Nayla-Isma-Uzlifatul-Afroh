package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return Celcius * 4 / 5
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	return (Celcius * 9 / 5) + 32
}

func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&c)

	fmt.Printf("\n%g celcius = %g reamur\n", c, CelciusToReamur(c))
	fmt.Printf("%g celcius = %g fahrenheit\n", c, CelciusToFahrenheit(c))
	fmt.Printf("%g celcius = %g kelvin\n", c, CelciusToKelvin(c))
}