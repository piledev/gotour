// // 実験したいこと： S を return すると、m() が実行されるのかどうか
// // I インターフェース を errorインターフェースと見立ててやってみたがどうにもうまく行かない
// // error インターフェースは特別扱いなのかもしれない

// package main

// import "fmt"

// // I :
// type I interface {
// 	M() string
// }

// // S : I インターフェースに適合させようと思っている型
// type S float64

// // M : Error()にあたるもの。自動的に実行されてほしい
// func (s S) M() string {
// 	return fmt.Sprintf("SのM()が実行されたよ！: %v", s)
// }

// // Sqrtx : I インターフェースを返すメソッド
// func Sqrtx(x float64) (float64, I) {
// 	return 0, S(x)
// }

// func main() {
// 	// s := S("sだお")
// 	// fmt.Println(s.M())

// 	// M()は呼んでないけど実行されてほしい
// 	fmt.Println(Sqrtx(99))

// }
