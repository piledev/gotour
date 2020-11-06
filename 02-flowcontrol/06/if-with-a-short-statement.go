package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// v := ... は条件ではなくて、条件のための準備
	// v のスコープは if 文の中だけ
	// ちなみにmath.Pow(x,n) = xのn乗
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(math.Pow(3, 3))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
