package main

import "fmt"

func main() {
	var char rune
	var hasil int
	hasil = 0
	fmt.Scanf("%c", &char)
	for char != '#' {
		hasil = hasil + 1
		fmt.Scanf("%c", &char)
	}
	fmt.Println(hasil)
}
