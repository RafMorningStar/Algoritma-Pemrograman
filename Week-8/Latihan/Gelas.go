package main

import (
	"fmt"
)

func main() {
	var r, t int
	var LulasTa, LuLasKer, TingKer, VolKer, VolTa float64
	fmt.Scan(&r, &t)

	LuLasKer = (22.0 / 7.0) * float64(r*r)
	TingKer = (2.0 / 3.0) * float64(t)
	VolKer = (1.0 / 3.0) * (LuLasKer * TingKer)

	LulasTa = (22.0 / 7.0) * ((1.0 / 5.0) * float64(r)) * ((1.0 / 5.0) * float64(r))
	VolTa = LulasTa * (float64(t) - TingKer)

	fmt.Printf("%.3f", VolKer+VolTa)

}
