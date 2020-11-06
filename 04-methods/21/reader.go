package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 文字列のReaderをつくる。
	r := strings.NewReader("Hello, Reader!")
	// 8マスのbyteスライスを作る。
	b := make([]byte, 8)
	for {
		// rを読み込みbに格納する。nは入れたバイトのサイズ。
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		// スライスはクリアせず上書きされるので、
		// 切りよく終わらない限り前回の読み込み結果が残る。
		// Hello, R
		// eader! R ←これな

		if err == io.EOF {
			break
		}

	}
}
