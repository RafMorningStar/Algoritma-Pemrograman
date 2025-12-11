package main

import "fmt"

func main() {
	var bil, i int
	var isPrima bool
	fmt.Scan(&bil)

	isPrima = true
	if bil <= 1 {
		isPrima = false
	} else if bil == 2 {
		isPrima = true
	} else if bil%2 == 0 {
		isPrima = false
	} else {
		for i = 3; i*i <= bil; i += 2 {
			if bil%i == 0 {
				isPrima = false
				break
			}
		}
	}
	if isPrima {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
