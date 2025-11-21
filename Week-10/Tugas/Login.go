package main
import "fmt"

func main() {
	var usn, pass string
	fmt.Scan(&usn, &pass)

	if usn == "admin" && pass == "12345" {
		fmt.Println("Login Berhasil")
	} 

	if usn != "admin" && pass != "12345" {
		fmt.Printf("Username Tidak Ditemukan\nPassword salah\n")
	}

	if usn != "admin" && pass == "12345"{
		fmt.Println("Username Tidak Ditemukan")
	}

	if usn == "admin" && pass != "12345" {
		fmt.Println("Password salah")
	}
}