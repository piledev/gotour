package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 暗黙的な型宣言は := で表す
	// 関数の外では var 宣言が必要で、 := による暗黙的な宣言はできない
	k := 3
	c, python, java := true, false, "no!"
	// = var c, python, java = true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
