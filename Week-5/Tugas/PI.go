package main

import "fmt"

func main() {
	var n, i int
	var pecahan, sign, kurung float64
	fmt.Scan(&n)

	sign = 1
	kurung = 1
	for i = 1; i <= n; i++ {
		sign = sign * -1
		pecahan = (1 / (1 + (float64(i) * 2))) * sign
		kurung = kurung + pecahan
	}
	fmt.Println(4 * kurung)
}
