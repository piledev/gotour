package main

import "fmt"

func main() {
	names := [4]string{
		"0:John",
		"1:Paul",
		"2:George",
		"3:Ringo",
	}
	fmt.Println(names)

	// スライス(aとb)は配列への参照
	// なので、スライスを定義する際は必ず参照先を指定してやる必要がある
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスの要素を変更すると、参照先の配列の要素が変更される。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
