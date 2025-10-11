package main
import "fmt"

func main() {
	var n, b, d1, d2, total int
	fmt.Scan(&n)
	total = 0
	for i := 1; i <= n; i++{
		fmt.Scan(&b)
		d1 = b/1000
		d2 = b%10
		total = total + d1 + d2
	}
	fmt.Println(total)
}
