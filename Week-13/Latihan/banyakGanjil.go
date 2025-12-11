package main

import "fmt"

func main() {
	var n, banyak int
	fmt.Scan(&n)
	banyak = 0
	for n > 0 {
		if (n-1)%2 == 0 {
			banyak++
		}
		n--
	}
	fmt.Printf("Terdapat %d bilangan ganjil", banyak)
}
