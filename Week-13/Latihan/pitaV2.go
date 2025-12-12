package main

import "fmt"

func main() {
	var i int
	var bunga, pita string
	i = 1
	fmt.Printf("Bunga %d : ", i)
	fmt.Scan(&bunga)
		for bunga != "SELESAI" {
			pita = pita + bunga + " " + "-" + " "
			i++
			fmt.Printf("Bunga %d : ", i)
			fmt.Scan(&bunga)
		}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", i-1)
}
