package main

import "fmt"

func main() {
	var d, jam, menit, detik int
	fmt.Scan(&d)
	jam = d / 3600
	menit = (d % 3600) / 60
	detik = (d % 3600) % 60
	fmt.Printf("%d jam %d menit %d detik\n", jam, menit, detik)

}
