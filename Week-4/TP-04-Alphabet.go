package main

import "fmt"

func main() {
	var C rune
	var hasil bool
	fmt.Scanf("%c",&C)

	hasil = (C >= 'a' && C <= 'z') || (C >= 'A' && C <= 'Z')
	fmt.Println(hasil)
}
