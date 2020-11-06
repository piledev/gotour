package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// もし m に key があれば変数 ok は True となり、なければ Falese となる。
	// また、key がないときは v は map要素の型のゼロ値となる。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
