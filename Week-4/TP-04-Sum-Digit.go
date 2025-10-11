package main

import "fmt"

func main() {
	var A, B, d1, d2, d3, d4 int
	fmt.Scan(&A)

	d4 = A / 1000
	d3 = (A / 100) % 10
	d2 = (A / 10) % 10
	d1 = A % 10

	B = d4*1000 + d1*100 + d2*10 + d3

	fmt.Println(A + B)
}
