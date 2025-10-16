package main

import (
	"fmt"
	"math"
)

func main() {
	var baris, jari, tinggi, i int
	var volume float64
	fmt.Scan(&baris)

	for i = 1; i <= baris; i += 1 {
		fmt.Scan(&jari, &tinggi)
		volume = float64(jari*jari*tinggi) * 1/3 * math.Pi
		fmt.Println(volume)
	}

}
