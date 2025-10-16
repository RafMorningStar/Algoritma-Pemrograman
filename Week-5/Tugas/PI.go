package main
import "fmt"

func main() {
	var n, i int
	var pecahan, sign float64
	fmt.Scan(&n)
	
	sign = -1
	for i = 1; i <= n; i++ {
	
	pecahan = 1/(1+(float64(i)*2)) * sign
	}
	fmt.Println(4*(1+pecahan))
}