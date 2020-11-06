package main

import "fmt"

func main() {

	// s := []int{2, 3, 5, 7, 11, 13}
	test := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = test[:]
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// s (slice)はあくまでポインタ
	// 今、sの0番目の要素は配列testの2番目の要素を指すポインタが入っている状態
	s[0] = 99
	fmt.Println(test)
}
