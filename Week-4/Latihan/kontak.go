package main

import "fmt"

func main() {
	type Kontak struct {
		nama    string
		telepon string
	}

	var kontak1 Kontak
	fmt.Scan(&kontak1.nama, &kontak1.telepon)

	fmt.Println("Nama:", kontak1.nama)
	fmt.Println("Telepon:", kontak1.telepon)
}
