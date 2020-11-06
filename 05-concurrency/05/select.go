package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select は複数ある case のいずれかが準備できるようになるまでブロックし、
		// 準備ができた case を実行する。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// default:
			// 	fmt.Println(".")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 無名関数を goroutine で実行。
	// (recieve from c)の表示を10回繰り返して、最後に quit に 0 を送信。
	// c からの受信はこのタイミングではまだ発生しないから実行されないに等しい。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// fibonacci の実行。
	// c <- x の準備はできているからまずはそれが繰り返される。
	// すると goroutine 内の <-c がひとつずつできるようになり for が回る。
	// 10 回繰り返されると for を抜けて quit に送信される。
	// すると、fibonacci 内で c <- x はできないが、<- quit はできる状態になり、
	// "quit"が表示され、リターンされる。
	fibonacci(c, quit)
}
