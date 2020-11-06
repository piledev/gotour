package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs :メソッドというのは、クラスのメソッドのようなもので、
// 他の言語ではクラスの中に書くような内容だが、Goでは外に書いている
// というイメージかな？
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
