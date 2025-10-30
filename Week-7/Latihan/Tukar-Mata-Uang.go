package main

import "fmt"

func main() {
	var awal, dinar, dirh, fals, qirat int
	fmt.Scan(&awal)
	dinar = awal / 600
	dirh = (awal % 600) / 60
	fals = ((awal % 600) % 60) / 6
	qirat = ((awal % 600) % 60) % 6
	fmt.Printf("%d %d %d %d\n", dinar, dirh, fals, qirat)

}
