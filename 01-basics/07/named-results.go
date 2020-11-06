package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// 変数名を指定せずにreturn
	// これをnaked returnという
	// 長い関数で使うと可読性が下がってしまう
	return
}

func main() {
	fmt.Println(split(17))
}
