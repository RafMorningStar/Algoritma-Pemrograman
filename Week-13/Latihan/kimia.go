package main

import "fmt"

func main() {
	var i, benar int
	var tab1, tab2, tab3, tab4 string
	for i = 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&tab1, &tab2, &tab3, &tab4)
		if tab1 == "merah" && tab2 == "kuning" && tab3 == "hijau" && tab4 == "ungu" {
			benar++
		}
	}
	if benar == 5 {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
