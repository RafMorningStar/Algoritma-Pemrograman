package main
import "fmt"

func main() {
	var x, y, i, hasil int
	fmt.Scan(&x, &y)
	hasil = 0
	for i = x; i <= y; i++{
		 hasil = hasil + i
	}
	fmt.Println(hasil)
}
