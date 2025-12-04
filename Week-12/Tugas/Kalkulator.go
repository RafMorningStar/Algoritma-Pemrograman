package main

import "fmt"

func main() {
	var target, bilangan, total int
	fmt.Scan(&target)
	for total < target {
		fmt.Scan(&bilangan)
		total = total + bilangan
		fmt.Print(bilangan, " ")
	}
	fmt.Println()
	fmt.Println(total)

}
