package main
import "fmt"

func main() {
	var bil1, bil2, hasil int
	fmt.Scan(&bil1, &bil2)
	for bil1 - bil2 >= 0 {
		bil1 = bil1 - bil2
		hasil= hasil + 1
	}
	fmt.Println(hasil)

}