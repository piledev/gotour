package main

import (
	"fmt"
	"math"
)

// Sqrt :exercise
func Sqrt(x float64) float64 {
	z := float64(1)
	pz := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)

		if (pz*pz-x)*(pz*pz-x) > (z*z-x)*(z*z-x) || ((pz-z)*(pz-z) > 1) {
			pz = z
		} else {
			fmt.Println("------------")
			return pz
		}
	}
	fmt.Println("------------")
	return z
}

func main() {
	x := float64(2)
	y := Sqrt(x)
	a := math.Sqrt(x)
	fmt.Printf("my answer    : %v\n", y)
	fmt.Printf("math's answer: %v\n", a)
	fmt.Println("------------")
	fmt.Println((a*a-2)*(a*a-2) >= (y*y-2)*(y*y-2))
}
