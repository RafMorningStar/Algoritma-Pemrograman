package main
import "fmt"

func main() {
	var kiri, kanan float64
	var oleng bool
	for kiri + kanan <= 150 || kiri < 0 || kanan < 0  {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
		oleng = kiri - kanan >= 9 || kanan - kiri >= 9
		if kiri + kanan <= 150 || kiri < 0 || kanan < 0 {
		fmt.Printf("Sepeda motor pak Andi akan oleng: %v\n", oleng)
		}
	}
		fmt.Println("Proses selesai.")
}
