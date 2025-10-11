package main
import "fmt"

func main(){
	var n, x, hasil int
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		fmt.Scan(&x)
		hasil = x*i 
		fmt.Print(hasil, " ")
	}
	fmt.Println()
}