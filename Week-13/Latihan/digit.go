package main

import "fmt"

func main() {
	var bil, i int
	fmt.Scan(&bil)
	i = 1
	for bil > 10 {
		bil = bil / 10
		i++
	}
	fmt.Println(i)

}
