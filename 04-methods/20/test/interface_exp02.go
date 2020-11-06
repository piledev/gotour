// 実験したいこと： Errorr を return すると、Errorr() が実行されるのかどうか
// exercise-errors の error 参照部分をすげ替えただけなのに自動的に実行はされない。
// error インターフェースは特別扱いなのだ。
// ではどうやって特別な挙動を実装しているのか？
package main

import (
	"fmt"
)

// ErrNegativeSqrt :
type ErrNegativeSqrt float64

// Errorr :
func (e ErrNegativeSqrt) Errorr() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

type Errorr interface {
	Errorr() string
}

// Sqrt :
func Sqrt(x float64) (float64, Errorr) {
	///////////////////////////////////
	if x < 0 {
		// ・ErrNegativeSqrt 型は Error() メソッドを持っているから error インターフェースに適合する
		// ・error インターフェースに適合するから error 型として返せる
		// ・error インターフェースを返すとなぜ Error() の実行結果が返るのか？
		//  仮説）error 型を返すと自動的に Error() が実行されるようになっている。
		return 0, ErrNegativeSqrt(x)
	}
	///////////////////////////////////
	z := float64(1)
	pz := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if (pz*pz-x)*(pz*pz-x) > (z*z-x)*(z*z-x) || ((pz-z)*(pz-z) > 1) {
			pz = z
		} else {
			return pz, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))

	_, b := Sqrt(-4)
	fmt.Printf("%T: %v\n", b, b)
}
