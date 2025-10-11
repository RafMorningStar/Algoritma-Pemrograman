package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 1; i <= n; i++ {
		fmt.Println(s)
	}

}
