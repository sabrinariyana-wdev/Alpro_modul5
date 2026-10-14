package main
import "fmt"
func main() {

	type buku struct {
		judulBuku string
		genre string
		tahunRilis string
		penilaian string
	}

	var b1, b2 buku

	fmt.Scan(&b1.judulBuku, &b2.judulBuku)
	fmt.Scan(&b1.genre, &b2.genre)
	fmt.Scan(&b1.tahunRilis, &b2.tahunRilis)
	fmt.Scan(&b1.penilaian, &b2.penilaian)


	fmt.Println("--- Book 1 ---")
	fmt.Println("Judul Buku:", b1.judulBuku)
	fmt.Println("Genre:", b1.genre)
	fmt.Println("Tahun Rilis:", b1.tahunRilis)
	fmt.Println("Penilaian:", b1.penilaian)

	fmt.Println("--- Book 2 ---")
	fmt.Println("Judul Buku:", b2.judulBuku)
	fmt.Println("Genre:", b2.genre)
	fmt.Println("Tahun Rilis:", b2.tahunRilis)
	fmt.Println("Penilaian:", b2.penilaian)
}

// program Buku
// kamus
// 	type buku <
// 	judulBuku : string
// 	genre : string
// 	tahunRilis : string
// 	penilaian : string >

// 	b1, b2 : buku
// algoritma
// 	input(b1.judulBuku, b2.judulBuku, b1.genre, b2.genre, b1.tahunRilis, b2.tahunRilis, b1.penilaian, b2.penilaian)
// 	output(b1.judulBuku, b2.judulBuku, b1.genre, b2.genre, b1.tahunRilis, b2.tahunRilis, b1.penilaian, b2.penilaian)
// endprogram