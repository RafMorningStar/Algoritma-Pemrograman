package main

import "fmt"

func main() {
	var selisih, huruf, hasil rune
	fmt.Scanf("%c", &huruf)
	selisih = 'a' - 'A'
	hasil = huruf + selisih
	fmt.Printf("%c\n", hasil)
}
