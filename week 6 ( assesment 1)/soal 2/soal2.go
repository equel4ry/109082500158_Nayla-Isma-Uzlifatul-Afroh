package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari(jumlahMhs int) int {
	makan := 2 * 35000
	penginapan := 250000
	uangSaku := 300000

	totalPerOrang := makan + penginapan + uangSaku
	return totalPerOrang * jumlahMhs
}

func perhitunganBiaya(jumlahMhs int, lamaPerjalanan int, tujuan string, totalBiaya *float64) {

	// hitung hari yang ditanggung
	hari := tanggunganHari(lamaPerjalanan, tujuan)

	// biaya harian domestik
	totalHarian := biayaPerHari(jumlahMhs)

	// jika mancanegara, dikali 1.5
	if tujuan == "mancanegara" {
		totalHarian = int(float64(totalHarian) * 1.5)
	}

	// total biaya akhir
	*totalBiaya = float64(totalHarian * hari)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	// input
	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	// proses perhitungan
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	// output
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}