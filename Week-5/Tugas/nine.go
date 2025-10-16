package main
import "fmt"

func main() {
	var i, n, bil, d1, d2, hasil int
	fmt.Scan(&n)
	hasil = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		d1 = (bil / 10) / 9
		d2 = (bil % 10) / 9
		
		hasil = hasil + d1 + d2
		
	}
	fmt.Println(hasil)
}