// 自力でできた！
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	ch := make(chan int, 10)

	fmt.Println("t1t1t1t1t1t1t1t1")
	go Walk(t1, ch)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("t2t2t2t2t2t2t2t2")
	go Walk(t2, ch)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("t1 == t2 ? :", Same(t1, t2))
}
