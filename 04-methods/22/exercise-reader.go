// 自力でできませんでした！
package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

// MyReader :
type MyReader struct{}

// 問題が何を言っているかわからなかったので答えを書き写しただけ。
// https://gist.github.com/CarlosLanderas/11b4f6727deec051883ddc02edf5cd0b

// 問題： ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。
// 解説： 値を入れたいスライスを渡すとすべて'A'で埋め尽くしてくれるReaderを実装してください。
// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})

	// 何をやっているのか確認する。
	var r MyReader
	b := make([]byte, 10)

	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:])

}
