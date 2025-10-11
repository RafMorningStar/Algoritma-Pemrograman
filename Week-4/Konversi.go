package main

import "fmt"

func main() {
	var x, selisih, hasil rune
	fmt.Scan(&x)
	selisih = 'a' - 'A'
	hasil = x - selisih
	fmt.Printf("%c\n", hasil)

}
