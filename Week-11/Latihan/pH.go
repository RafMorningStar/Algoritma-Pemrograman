package main

import "fmt"

func main() {
	var pH float64
	fmt.Scan(&pH)
	switch {
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air layak minum")
	case (pH >= 0 && pH < 6.5) || (pH > 8.6 && pH <= 14):
		fmt.Println("Air tidak layak minum")
	default:
		fmt.Println("Nilai PH tidak valid. Nilai pH harus antara 0 dan 14")
	}

}
