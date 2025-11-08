package main

import "fmt"

func main() {
	var emas, perak, tembaga, timah, sisatimah int
	fmt.Scan(&timah)
	emas = timah / 600
	perak = (timah % 600) / 60
	tembaga = (timah % 60) / 6
	sisatimah = timah % 6

	fmt.Println(emas, perak, tembaga, sisatimah)
}
