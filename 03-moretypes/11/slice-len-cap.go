package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s)
	// len=4 cap=6 [2 3 5 7]
	// Q.
	// 一旦len=0になったものがなぜ復活した？
	// A.
	// sliceするとき（[:]などを使って範囲指定する場合）は
	// lengthではなくcapacityの範囲は参照できる。
	// つまり
	// sliceの最初の要素〜配列の最後の要素までは生き続けているということ。

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	// len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	// len() は要素数
	// cap() は容量（スライスの最初の要素から数えて元となる配列の要素数）
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
