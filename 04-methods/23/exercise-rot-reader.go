// 自力でできませんでした！
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return
}

func main() {
	// 文字列のReaderを作る
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	// rot13Reader 構造体の r の値に *strings.Reader を格納し、main.r に代入する。
	// (*strings.Reader は io.Reader に該当するので rot13Reader に格納できる）
	r := rot13Reader{s}

	// io.Copy : &r を os.Stdout （標準出力）にコピーする（＝表示する）
	// おそらくコピーの処理の中で io.Reader の Read が呼び出されるようになっているから Read が実行される
	io.Copy(os.Stdout, &r)
}
