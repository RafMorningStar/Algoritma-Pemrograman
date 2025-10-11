package main
import "fmt"

func main () {
	var hari, jam, totjam int
	var rata float64
	fmt.Scan(&hari)
	totjam = 0
	for i := 1; i <= hari; i++{
		fmt.Scan(&jam)
		totjam = totjam + jam
	}
	rata = float64(totjam)/float64(hari)
	fmt.Printf("%.3f\n", rata)
}