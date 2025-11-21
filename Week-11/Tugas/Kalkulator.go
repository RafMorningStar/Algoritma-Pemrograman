package main

import "fmt"

func main() {
	var ops, bil1, bil2 int
	fmt.Scan(&bil1, &bil2, &ops)
	switch ops {
	case 1:
		fmt.Println(bil1 + bil2)
	case 2:
		fmt.Println(bil1 - bil2)
	case 3:
		fmt.Println(bil1 * bil2)
	case 4:
		if bil2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(float64(bil1) / float64(bil2))
		}
	case 5:
		if bil2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(bil1 % bil2)
		}
	case 6:
		if bil2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(bil1 / bil2)
		}
	}

}
