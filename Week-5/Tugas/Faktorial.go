package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i >= 2; i-- {
		fmt.Print(i, " x ")
	}
	fmt.Println(1)
}