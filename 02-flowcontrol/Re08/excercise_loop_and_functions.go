package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	// for i := 0; i < 10; i++ {
	// 	z -= (z*z - x) / (2 * z)
	// 	fmt.Println(z)
	// }
	for d := 1.0; d*d > 1e-10; z -= d {
		d = (z*z - x) / (2 * z)
		fmt.Println(d)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
