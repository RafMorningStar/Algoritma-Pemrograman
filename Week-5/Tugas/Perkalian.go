package main
import "fmt"

func main(){
	var N, b1, b2 int
	fmt.Scan(&N)
	for i := 1; i <= N; i++{
		fmt.Scan(&b1, &b2)
		fmt.Println(b1*b2)
	}
}