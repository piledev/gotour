// closure とは「外部の変数を参照する関数値」と定義される。
// この例でいうと adder() の中の func(x int) int が closure。
// この closure は adder() スコープの変数 sum にアクセスすることができる。
// オブジェクト指向で言うところのインスタンス内の変数にアクセスする振る舞い（関数）といった感じ？
package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// pos と neg には別々の adder() が割り当てられている。
	// 当然、それぞれの中の変数 sum は互いに独立している。
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
