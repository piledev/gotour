package main

import "fmt"

func main() {
	fmt.Println("counting")

	// defer へ渡した関数が複数ある場合、その呼出はスタックされる。
	// Defer is commonly used to simplify functions that perform various clean-up actions.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
