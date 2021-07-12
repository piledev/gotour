package main

import "fmt"

func main() {
	i, j := 42, 2701

	// ポインタとは値型に分類されるデータ構造のメモリ上のアドレスと型の情報。
	// ポインタ型は *int のように、ポインタを使って参照・操作する型の前に * を置くことで定義できる。
	// また、アドレス演算子 & を使って、任意の型からそのポインタ型を生成することができる。
	// 演算子 * をポインタ型の変数の前に置くことで、ポインタ型が保持するメモリ上のアドレスを経由してデータ本体を参照することができる。

	// &オペレータはその変数へのポインタ（＝アドレスや型の情報）
	p := &i // point to i

	// *オペレータはそのポインタが指す先の値
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	// こんな表現もできる（意味あるかどうかは別として）
	fmt.Println(*&i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
