package main

import "fmt"

const Pi = 3.14

// const(定数)は:=は使えない
// 文字, 文字列, boolean, 数値のみで使える

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
