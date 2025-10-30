package main

import "fmt"

func main() {
	var p, l, t, volume, luasp float64
	fmt.Scan(&p)
	fmt.Scan(&l)
	fmt.Scan(&t)
	volume = (p * l) * t
	luasp = 2*(p*l) + 2*(l*t) + 2*(p*t)
	fmt.Printf("%.1f\n%.1f\n", volume, luasp)
}
