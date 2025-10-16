package main

import "fmt"

func main() {
	var bil, i, hasil int
	fmt.Scan(&bil)
	hasil = 0
	for i = 1; i <= bil; i += 1 {
		hasil = hasil + i
	}
	fmt.Println(hasil)
}
