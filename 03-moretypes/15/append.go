package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// my original test
	a := []int{111, 222, 333} // array
	b := a[1:2]               // slice
	printSlices(a, b)
	// array:len=3 cap=3 [111 222 333]
	// slice:len=1 cap=2 [222]

	b = append(b, 999)
	printSlices(a, b)
	// array:len=3 cap=3 [111 222 999]
	// slice:len=2 cap=2 [222 999]
	// slice へ値を追加すると、参照先の array の該当位置の値が更新される。

	b = append(b, 999)
	printSlices(a, b)
	// array:len=3 cap=3 [111 222 999]
	// slice:len=3 cap=4 [222 999 999]
	// slice は array の capacity を超えて値を追加したので、slice の参照先は別の array に切り替わった。

	b[0] = 0
	printSlices(a, b)
	// array:len=3 cap=3 [111 222 999]
	// slice:len=3 cap=4 [0 999 999]
	// 元の array はもう参照先ではないため、slice の内容を更新しても元の array には影響しない。
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlices(a []int, b []int) {
	fmt.Printf("array:len=%d cap=%d %v\n", len(a), cap(a), a)
	fmt.Printf("slice:len=%d cap=%d %v\n", len(b), cap(b), b)
}
