package main

import "fmt"

func main() {
	// array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice : 定義上配列の長さを指定していないだけ
	var s []int = primes[1:4]
	fmt.Println(s)
}
