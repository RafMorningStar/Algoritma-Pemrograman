package main

import "fmt"

func main() {
	var ajuHari string
	var totalHari, selisihTanggal int
	var kabisat bool
	var ajuTanggal, ajuBulan, ajuTahun, ambilTanggal, ambilBulan, ambilTahun int
	fmt.Scan(&ajuHari)
	kabisat = true
	selisihTanggal = 2
	for ajuHari != "Exit" {
		fmt.Scan(&ajuTanggal, &ajuBulan, &ajuTahun)
		if (ajuTahun%400 == 0) || (ajuTahun%4 == 0 && ajuTahun%100 != 0) {
			kabisat = true
		} else {
			kabisat = false
		}
		if kabisat && ajuBulan == 2 {
			totalHari = 29
		} else if !kabisat && ajuBulan == 2 {
			totalHari = 28
		} else if ajuBulan == 4 || ajuBulan == 6 || ajuBulan == 9 || ajuBulan == 11 {
			totalHari = 30
		} else {
			totalHari = 31
		}
		if ajuHari == "Kamis" || ajuHari == "Jumat" {
			selisihTanggal = 4
		}
		if ajuTanggal+selisihTanggal > totalHari && ajuBulan == 12 {
			ambilTanggal = (ajuTanggal + selisihTanggal) - 31
			ambilBulan = 1
			ambilTahun = ajuTahun + 1
		} else if ajuTanggal+selisihTanggal > totalHari {
			ambilTanggal = (ajuTanggal + selisihTanggal) - totalHari
			ambilBulan = ajuBulan + 1
			ambilTahun = ajuTahun
		} else {
			ambilTanggal = ajuTanggal + selisihTanggal
			ambilBulan = ajuBulan
			ambilTahun = ajuTahun
		}
		fmt.Printf("passport bisa diambil pada tanggal %d bulan %d tahun %d\n", ambilTanggal, ambilBulan, ambilTahun)
		selisihTanggal = 2
		fmt.Scan(&ajuHari)
	}

}
