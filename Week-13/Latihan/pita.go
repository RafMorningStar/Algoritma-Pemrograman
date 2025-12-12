package main
import  "fmt"

func main() {
	var N, i int
	var bunga, pita string
	fmt.Print("N: ")
	fmt.Scan(&N)
	
	for i = 1; i <= N; i++ {
		fmt.Printf("Bunga %d : ", i)
		fmt.Scan(&bunga)
		pita = pita + bunga + " " + "-" + " "
	}
	fmt.Println("Pita: ", pita)
}