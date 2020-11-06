package main

import "fmt"

type T int

func (t *T) M() {
	fmt.Print(int(*t))
}

func main() {
	ts := []T{10, 20, 30}

	ms := make([]func(), 0, len(ts))

	for _, t := range ts {

		fmt.Println(t, &t)

		// 自分の解釈：
		// 変数 t に設定される値は for の繰り返しの中で変わっているが、
		// 変数 t 自身のポインタ（アドレス）はずっと変わらない。
		// ms スライスのすべての要素に t ポインタが設定される。
		// ということは、どの要素も t の値を指していて一致する。
		// t ポインタに最後に設定されるのは 30 なので、
		// どの要素も最終的な値は 30 になる。

		// ms = append(ms, (&t).M)
		// ↓のように書いても Go は↑のように解釈する。
		ms = append(ms, t.M)
	}

	// 303030と表示される
	for _, m := range ms {
		m()
	}
}
