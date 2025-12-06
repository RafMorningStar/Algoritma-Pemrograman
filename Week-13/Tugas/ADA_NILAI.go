package main

import "fmt"

func main() {
	var dicari, bilangan, posisi int
	var ada bool
	fmt.Scan(&dicari)
	fmt.Scan(&bilangan)
	ada = false
	posisi = 1
	for bilangan != 0 {
		if dicari == bilangan && !ada {
			ada = true
			fmt.Println(posisi)
		}
		posisi++
		fmt.Scan(&bilangan)
	}
	if !ada {
		fmt.Println("TIDAK ADA")
	}
}
