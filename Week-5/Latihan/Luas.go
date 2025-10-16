package main

import "fmt"

func main() {
	var n, i, palas, tinggi int
	var hasil float64
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&palas, &tinggi)
		hasil = 0.5 * float64(palas*tinggi)
		fmt.Println(hasil)
	}
}
