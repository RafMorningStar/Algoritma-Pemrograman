package main
import "fmt"

func main() {
	var orang, langkah, total int
	fmt.Scan(&orang)
	for i := 1; i <= orang; i++ {
		total = 0
		for j := 1; j <= 7; j++ {
			fmt.Scan(&langkah)
			total = total + langkah
		}
		fmt.Println(total)
	}
}
