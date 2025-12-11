package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, donaturBerapa int
	fmt.Scan(&target)
	donaturBerapa = 0
	totalDonasi = 0
	for totalDonasi < target {
		fmt.Scan(&donasi)
		donaturBerapa++
		totalDonasi = totalDonasi + donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donaturBerapa, donasi, totalDonasi)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur", totalDonasi, donaturBerapa)
}
