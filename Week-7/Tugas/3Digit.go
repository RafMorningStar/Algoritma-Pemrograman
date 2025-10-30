package main

import "fmt"

func main() {
	var d1, d2, d3, dd int
	fmt.Scanf("%1d%1d%1d\n", &d1, &d2, &d3)
	dd = ((d1*10 + d1) * 10000) + ((d2*10 + d2) * 100) + (d3*10 + d3)
	fmt.Println(dd)
	fmt.Println(dd * dd)

}
