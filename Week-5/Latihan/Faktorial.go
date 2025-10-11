package main
import "fmt"

func main() {
	var n, total int
	fmt.Scan(&n)
	total = 1
	for i := n; i >= 1; i-- {
		total = total * i
	}
	fmt.Println(total)
}