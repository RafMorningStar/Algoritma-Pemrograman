package main

import "fmt"

func main() {
	var total1, total2 int
	var kar rune

	fmt.Scanf("%c", &kar)

	for kar != '.' {
		if kar == 'A' || kar == 'a' || kar == 'I' || kar == 'i' || kar == 'U' || kar == 'u' ||
			kar == 'E' || kar == 'e' || kar == 'O' || kar == 'o' {
			total1 = total1 + 1
		} else if kar >= 'a' && kar <= 'z' || kar >= 'A' && kar <= 'Z' {
			total2 = total2 + 1
		}
		fmt.Scanf("%c", &kar)
	}
	fmt.Printf("Banyaknya huruf vokal: %v\nBanyaknya huruf konsonan: %v", total1, total2)
}
