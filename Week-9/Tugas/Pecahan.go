package main

import "fmt"

func main() {
	var hutangA, hutangB, lembarA, lembarB int
	fmt.Scan(&hutangA, &hutangB)

	lembarA = hutangA / 2000
	if hutangA%2000 > 0 {
		lembarA = lembarA + 1
	}
	lembarB = hutangB / 2000
	if hutangB%2000 > 0 {
		lembarB = lembarB + 1
	}
	fmt.Printf("PT. Alice memerlukan %d lembar USD 2,000\nPT. Bob memerlukan %d lembar USD 2,000\n", lembarA, lembarB)
}
