package main

import "fmt"

func main() {
	var x rune
	fmt.Scanf("%c", &x)
	if x-122 >= -25 && x-122 <= 0 {
		fmt.Println("Huruf kecil")
	} else {
		if x-90 >= -25 && x-90 <= 0 {
			fmt.Println("Huruf Besar")
		} else {
			if x-47 >= -14 && x-47 <= 0 {
				fmt.Println("Simbol")
			} else {
				fmt.Println("Bilangan")
			}
		}

	}

}
