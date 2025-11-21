package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	if x - 12 <= 0 {
		fmt.Println("Anak-anak")
	} else {
		if x - 17 >= -4 && x - 17 <=0 {
			fmt.Println("Remaja")
 		} else {
			if x - 55 >= -37 && x - 55 <= 0 {
				fmt.Println("Dewasa")
			} else {
				fmt.Println("Lansia")
			}
		}
	}
}
