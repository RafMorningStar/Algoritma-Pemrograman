package main

import "fmt"

func main() {
	var bil, pangkat, hasil, i int
	fmt.Scan(&bil, &pangkat)
	hasil = 1
	for i = 1; i <= pangkat; i += 1 {
		hasil = hasil * bil

	}
	fmt.Println(hasil)
}
