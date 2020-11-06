package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapの定義
// key は string, value は Vertex という意味だろう。
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// マップのゼロ値は nil です。nil マップはキーを持っておらず、またキーを追加することもできません。
	// という解説の意味がよくわからなかった。
}
