package main

import "fmt"

func main() {
	var N, i int
	fmt.Scan(&N)
	for i = 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}
