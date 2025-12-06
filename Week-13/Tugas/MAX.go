package main

import "fmt"

func main() {
	var bil1, bil2, max int
	fmt.Scan(&bil1, &bil2)
	for bil2 != 0 {
		if bil1 >= bil2 {
			max = bil1
		} else {
			bil1 = bil2
			max = bil2
		}
		fmt.Scan(&bil2)
	}
	fmt.Println(max)
}
