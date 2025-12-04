package main

import "fmt"

func main() {
	var bil, hasil int
	fmt.Scan(&bil)
	hasil = 0
	for bil > 1 {
		hasil = hasil + ((((bil % 10) % 2) - 1) * -1)
		bil = bil / 10
	}
	fmt.Println(hasil)

}
