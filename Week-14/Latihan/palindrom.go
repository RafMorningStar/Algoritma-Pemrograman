package main

import "fmt"

func main() {
	var kata string
	var palindrom bool = true
	var i, j int

	fmt.Scan(&kata)

	// i bergerak dari depan, j bergerak dari belakang
	// Loop berjalan selama i masih di sebelah kiri j
	for i, j = 0, len(kata)-1; i < j; i, j = i+1, j-1 {
		if kata[i] != kata[j] {
			palindrom = false
			break // Jika ada satu saja yang beda, langsung berhenti
		}
	}

	fmt.Println(palindrom)
}
