package main

import "fmt"

func main() {
	var bil float64
	var target int
	fmt.Scan(&bil)
	target = int(bil) + 1
	for bil < float64(target) {
		bil = bil + 0.1
		if bil < float64(target) && bil != float64(target) {
			fmt.Printf("%.1f\n", bil)
		} else {
			fmt.Println(target)
		}
	}
}
