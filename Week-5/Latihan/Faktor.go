package main
import "fmt"

func main() {
	var N, d int
	var s bool
	fmt.Scan(&N)
	for i := 1; i <= N; i++{
		d = i
		s = N % i == 0
		fmt.Println(d, s)
	}
}