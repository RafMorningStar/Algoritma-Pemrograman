package main
import "fmt"

func main() {
	var jumlahGagal int
	var username, password string
	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin"{
		jumlahGagal = jumlahGagal + 1
		fmt.Scan(&username, &password)
	}
	fmt.Printf("%d percobaan gagal login\n", jumlahGagal)
}