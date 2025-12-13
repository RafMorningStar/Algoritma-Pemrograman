package main

import "fmt"

func main() {
	var nBuku int
	var totalhargaBuku float64
	var member string
	fmt.Scan(&nBuku, &member)
	totalhargaBuku = 10000 * float64(nBuku)
	switch member {
	case "A":
		if nBuku < 5 {
			totalhargaBuku = totalhargaBuku * 0.9
		} else if nBuku >= 5 && nBuku <= 10 {
			totalhargaBuku = totalhargaBuku * 0.8
		} else {
			totalhargaBuku = totalhargaBuku * 0.7
		}
	case "B":
		if nBuku < 5 {
			totalhargaBuku = totalhargaBuku * 0.95
		} else if nBuku >= 5 && nBuku <= 10 {
			totalhargaBuku = totalhargaBuku * 0.9
		} else {
			totalhargaBuku = totalhargaBuku * 0.85
		}
	case "C":
		if nBuku >= 5 && nBuku <= 10 {
			totalhargaBuku = totalhargaBuku * 0.95
		} else if nBuku > 10 {
			totalhargaBuku = totalhargaBuku * 0.9
		}
	case "N":
		if nBuku > 10 {
			totalhargaBuku = totalhargaBuku * 0.95
		}
	}
	fmt.Printf("Rp %.0f", totalhargaBuku)
}
