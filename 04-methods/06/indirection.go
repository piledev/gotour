package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	// 関数は引数の定義に合わせてポインタを渡す必要があるのに対して、
	// メソッドはポインタを渡しても変数を渡してもどちらでも良い。
	// Go が利便性のために気を利かせて
	// v.Scale(2) のステートメントを (&v).Scale(2) に読み替えてくれるから。
}
