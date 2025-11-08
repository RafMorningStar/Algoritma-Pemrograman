package main

import "fmt"

func main() {
	var selisih, hk, hb rune
	fmt.Scanf("%c", &hk)
	selisih = 'a' - 'A'
	hb = hk - selisih
	fmt.Printf("%c\n", hb)
}
