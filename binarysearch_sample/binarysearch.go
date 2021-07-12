package main

import (
	"fmt"
)

func main() {
	arr := [12]int{1, 3, 5, 6, 8, 10, 13, 20, 33, 45, 56, 89}
	high := len(arr)
	low := 0
	target := 33
	middle := (high + low) / 2
	i := 0
	for ; i < 10; i++ {
		fmt.Println(low, high, middle, arr[middle])
		switch {
		case arr[middle] > target:
			high = middle
		case arr[middle] < target:
			low = middle
		default:
			fmt.Println("Hit!")
			goto L
		}
		middle = (high + low) / 2
	}
L:
	fmt.Println(low, high, middle, arr[middle])

}
