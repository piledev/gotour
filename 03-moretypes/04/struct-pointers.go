package main

import "fmt"

// Vertex : struct
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// (*p).X と書くこともできるが、面倒すぎるので
	// go は p.X という省略記法も用意してくれた。
	p.X = 1e9
	fmt.Println(v)
}
