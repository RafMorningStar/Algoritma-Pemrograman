package main

import "fmt"

func main() {
	var jumsis, jumbus int
	fmt.Scan(&jumsis)
	jumbus = jumsis / 45
	if jumsis % 45 > 0{
		jumbus = jumbus + 1
	}
	fmt.Printf("Diperlukan %d bus untuk tamasya ke Lembang\n", jumbus)
}
