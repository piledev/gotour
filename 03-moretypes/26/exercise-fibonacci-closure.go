package main

import "fmt"

// 20210515 に書いたコード
// func fibonacci() func() int {
// 	count := 0
// 	s := []int{0,1}
// 	fn := func() int {
// 		count++
// 		switch count {
// 			case 1,2:
// 				return s[count-1]
// 			default:
// 				s = append(s, s[count-3] + s[count-2])
// 				return s[len(s)-1]
// 		}
// 	}
// 	return fn
// }

func fibonacci() func() int {
	f := [3]int{0, 0, 1}
	ret := 0
	return func() int {
		ret = f[1]
		f[0] = f[1]
		f[1] = f[2]
		f[2] = f[0] + f[1]
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
