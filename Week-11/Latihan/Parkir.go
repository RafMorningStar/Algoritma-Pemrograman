package main

import "fmt"

func main() {
	var kendaraan string
	var durasiParkir int
	fmt.Scan(&kendaraan, &durasiParkir)
	switch kendaraan {
	case "Motor":
		if durasiParkir < 2 || durasiParkir == 2 {
			fmt.Println("Tarif Parkir: Rp 7000")
		} else {
			fmt.Println("Tarif Parkir: Rp 9000")
		}
	case "Mobil":
		if durasiParkir < 2 || durasiParkir == 2 {
			fmt.Println("Tarif Parkir: Rp 15000")
		} else {
			fmt.Println("Tarif Parkir: Rp 20000")
		}
	case "Truk":
		if durasiParkir < 2 || durasiParkir == 2 {
			fmt.Println("Tarif Parkir: Rp 25000")
		} else {
			fmt.Println("Tarif Parkir: Rp 30000")
		}
	default:
		fmt.Println("Jenis Kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
	}
}
