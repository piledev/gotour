package main

import "fmt"

func fibonacci() func() int {
	f := [3]int{0, 0, 1}
	ret := 0
	return func() int {
		ret = f[1]
		f[0] = f[1]
		f[1] = f[2]
		f[2] = f[0] + f[1]
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
