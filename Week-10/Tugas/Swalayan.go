package main

import "fmt"

func main() {
	var membership string
	var a, b, c, d, e int
	var cashback, diskon float64
	fmt.Scan(&membership, &a, &b, &c, &d, &e)

	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 3.1 * float64(a+b+c)
		diskon = 0
	} else {
		if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
			cashback = 0
			diskon = 1.7 * float64(c+d+e)
		} else {
			cashback = 3.1 * float64(a+b+c)
			diskon = 1.7 * float64(c+d+e)
		}
	}
	if membership == "yes" {
		diskon = diskon * 1.15
		cashback = cashback * 1.15

	}
	if cashback > 35.0 {
		cashback = 35.0
	}
	if diskon > 50.0 {
		diskon = 50.0
	}
	fmt.Printf("cashback %.2f, dan discount %.2f\n", cashback, diskon)
}
