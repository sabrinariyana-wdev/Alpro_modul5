package main
import "fmt"

func main() {
	const diskon1 float64 = 0.10
	const diskon2 float64 = 0.20
	var hargaAwal, hargaSetelahDiskon1, hargaAkhir float64

	fmt.Scan(&hargaAwal)
	hargaSetelahDiskon1 = hargaAwal - (hargaAwal * diskon1)
	hargaAkhir = hargaSetelahDiskon1 - (hargaSetelahDiskon1 * diskon2)

	fmt.Printf("%.2f", hargaAkhir)
}