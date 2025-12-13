package main

import "fmt"

func main() {
	var WPM1, WPM2, WPM3, totalWPM, averageWPM float64
	var set int
	for (WPM1 < 90.0 && WPM2 < 90.0 && WPM3 < 90.0) || totalWPM < 210.0 {
		fmt.Scan(&WPM1, &WPM2, &WPM3)
		totalWPM = totalWPM + WPM1 + WPM2 + WPM3
		set++
	}
	averageWPM = totalWPM / (3.0 * float64(set))
	fmt.Printf("Total set: %d (rata-rata: %.2f)", set, averageWPM)

}
