package main
import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	for bilangan > 0 {
		fmt.Println(bilangan % 10)
		bilangan = bilangan / 10
	}

}