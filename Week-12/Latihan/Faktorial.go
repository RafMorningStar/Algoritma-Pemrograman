package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	for bil > 1 {
		fmt.Print(bil, " x ")
		bil = bil - 1
	}
	fmt.Println("1")
}
