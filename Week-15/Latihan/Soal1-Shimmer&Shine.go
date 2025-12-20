package main

import "fmt"

func main() {
	var intens1, intens2 int
	fmt.Scan(&intens1, &intens2)
	if (intens1%2 == 0 && intens2%2-1 == 0) || (intens1%2-1 == 0 && intens2%2 == 0) {
		fmt.Println("berhasil")
	}

}
