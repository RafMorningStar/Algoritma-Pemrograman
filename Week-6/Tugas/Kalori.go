package main

import "fmt"

func main() {
	var set, pu, su, sq, jj int
	var i, j, k, l, total float64
	fmt.Scan(&set)
	fmt.Scan(&pu, &su, &sq, &jj)
	total = 0

	for i = 1; i <= float64(pu); i++ {
		total = total + i*0.5
	}
	for j = 1; j <= float64(su); j++ {
		total = total + j*0.3
	}
	for k = 1; k <= float64(sq); k++ {
		total = total + k*0.2
	}
	for l = 1; l <= float64(jj); l++ {
		total = total + l*0.6
	}

	fmt.Println("Total kalori yang terbakar hari ini sebanyak", float64(set)*total, "kalori")

}
