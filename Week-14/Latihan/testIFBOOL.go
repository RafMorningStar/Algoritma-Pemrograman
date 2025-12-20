package main

import "fmt"

func main() {
	var kabisat, ganteng bool
	fmt.Println(kabisat, ganteng)

	fmt.Scan(&kabisat)

	if kabisat {
		fmt.Println(ganteng)
	} else {
		fmt.Println(!ganteng)
	}

}
