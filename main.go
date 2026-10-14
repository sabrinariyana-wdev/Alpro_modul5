package main
import "fmt"

func main() {
	const(
		kuis float64 = 0.15
		tugas float64 = 0.15
		praktikum float64 = 0.2
		uts float64 = 0.25
		uas float64 = 0.25
	)
	type mahasiswa struct {
		nim string
		nama string
		nilaiTugas float64
		nilaiKuis float64
		nilaiUas float64
		nilaiPraktikum float64
		nilaiAkhir float64
	}
	var dataMahasiswa mahasiswa 

	fmt.Scan(&dataMahasiswa.nama, &dataMahasiswa.nim, &dataMahasiswa.nilaiTugas, 
		&dataMahasiswa.nilaiKuis, &dataMahasiswa.nilaiUas, &dataMahasiswa.nilaiPraktikum)

	dataMahasiswa.nilaiAkhir = dataMahasiswa.nilaiTugas*tugas + dataMahasiswa.nilaiKuis*kuis + dataMahasiswa.nilaiUas*uas + dataMahasiswa.nilaiPraktikum*praktikum
	
	fmt.Println("Data Mahasiswa :") 
	fmt.Println("Nama :", dataMahasiswa.nama)
	fmt.Println("NIM :", dataMahasiswa.nim)
	fmt.Println("Nilai Tugas:", dataMahasiswa.nilaiTugas)
	fmt.Println("Nilai Kuis :", dataMahasiswa.nilaiKuis)
	fmt.Println("Nilai uas :", dataMahasiswa.nilaiUas)
	fmt.Println("Nilai Praktikum :", dataMahasiswa.nilaiPraktikum)
	fmt.Println("Nilai Akhir :", dataMahasiswa.nilaiAkhir)
}