package main

import "fmt"

var i, j int = 1, 2

func main() {
	// var 宣言で初期化子が与えられている場合は型指定を省略できる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
