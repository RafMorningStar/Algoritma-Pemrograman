package main

import "fmt"

func main() {
	var bil1, bil2, i, hasil int
	fmt.Scan(&bil1, &bil2)
	hasil = 0
	for i = 1; i <= bil2; i += 1 {
		hasil = hasil + bil1

	}
	fmt.Println(hasil)
}
