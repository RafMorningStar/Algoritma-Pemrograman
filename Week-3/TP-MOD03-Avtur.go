package main

import "fmt"

func main() {
	var lb, kg, L float64
	fmt.Scan(&lb)

	kg = lb * 0.45359237
	L = kg / 0.80

	fmt.Printf("%.6f kg %.6f L\n", kg, L)

}
