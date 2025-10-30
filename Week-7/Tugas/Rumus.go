package main

import "fmt"

func main() {
	var x, n, i int
	var s float64
	fmt.Scan(&x, &n)
	n = n + 1
	s = 0
	for i = 1; i <= n-1; i++ {
		s = s + (float64(n)-float64(i))/(float64(i)*float64(x))
	}
	fmt.Printf("%.3f\n", s)

}
