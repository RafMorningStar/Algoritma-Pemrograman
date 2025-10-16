package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scan(&x, &y)

	for i = x; i <= y; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
