package main

import "fmt"

func main() {
	var bil1, bil2, bil3 int
	fmt.Scan(&bil1, &bil2, &bil3)
	if bil1 < 0 {
		bil1 = -bil1
	}
	if bil2 < 0 {
		bil2 = -bil2
	}
	if bil3 < 0 {
		bil3 = -bil3
	}
	fmt.Println(bil1, bil2, bil3)

}
