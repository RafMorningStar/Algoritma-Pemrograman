package main

import "fmt"

func main() {
	var Talice, Tbob int
	var hasil bool
	fmt.Scan(&Talice, &Tbob)
	hasil = Talice == Tbob || Talice-1 == Tbob || Talice+1 == Tbob || Talice-5 == Tbob || Talice+5 == Tbob
	fmt.Println("Menang?", hasil)
}
