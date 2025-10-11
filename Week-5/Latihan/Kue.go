package main

import "fmt"

func main() {
	var n, beratx, total int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&beratx)

		total = beratx * 10
		fmt.Println(total)
	}

}
