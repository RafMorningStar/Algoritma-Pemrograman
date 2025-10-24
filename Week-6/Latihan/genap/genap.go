package main
import "fmt"

func main() {
	var i, banyak, bilangan, hasil int
	fmt.Scan(&banyak)
	hasil = 0
	for i = 1; i <= banyak; i++ {
		fmt.Scan(&bilangan)
		hasil = hasil + ((bilangan%2-1)*-1)
	}
	fmt.Println(hasil)
}
