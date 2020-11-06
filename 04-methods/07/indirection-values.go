package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	// 関数は引数の定義に合わせて値を渡す必要があるのに対して、
	// メソッドはポインタを渡しても変数を渡してもどちらでも良い。
	// Go が利便性のために気を利かせて
	// p.Abs() のステートメントを (*p).Abs() に読み替えてくれるから。

}
