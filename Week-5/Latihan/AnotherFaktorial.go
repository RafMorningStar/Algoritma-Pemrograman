package main

import "fmt"

func main() {
	var i, bil, hasil int
	fmt.Scan(&bil)
	hasil = 1
	for i = bil; i >= 1; i -= 1 {
		hasil = hasil * i
	}
	fmt.Println(hasil)

}
