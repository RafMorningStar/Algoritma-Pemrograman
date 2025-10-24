package main

import "fmt"

func main() {
	var i, banyak, bilangan, hasil int
	fmt.Println("Input banyak bilangan")
	fmt.Scan(&banyak)
	hasil = 0
	for i = 1; i <= banyak; i++ {
		fmt.Println("Input bilangan ke-", i)
		fmt.Scan(&bilangan)
		hasil = hasil + bilangan
	}
	fmt.Println("hasil:", hasil)
}
