package main

import "fmt"

func main() {
	var id, ratusan, puluhan, satuan int
	fmt.Scan(&id)

	ratusan = id / 100
	puluhan = (id / 10) % 10
	satuan = id % 10

	fmt.Println(ratusan, puluhan, satuan)
}
