package main

import "fmt"

func main() {
	var bilangan, d2 int
	fmt.Scan(&bilangan)
	d2 = bilangan % 10
	switch {
	case d2 == 5 && bilangan != 5:
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^2 = %d\n", bilangan, bilangan*bilangan)
	case d2 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil kuadrat dari %d / 10 = %d\n", bilangan, bilangan/10)
	case bilangan%2 == 1:
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, bilangan+(bilangan+1))
	case bilangan%2 == 0:
		fmt.Printf("Kategori: Bilangan Genap\nHasil Perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, bilangan*(bilangan+1))
	}
}
