package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Println("hello world")
}

func swap(x, y string) (string, string) {
	return y, x
}

// cast
func changeType(i int) uint {
	f := float64(i)
	u := uint(f)
	return u
}

// for文
func For() int {
	// for文の基本
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// for文の省略パターン
	j := 1
	for j < 100 {
		j += j
	}
	fmt.Println(j)
	return sum
}

func If() {
	// if文の基本
	x := true
	if x {
		fmt.Println("x is true")
	}

	// if文で変数宣言するパターン
	if y := false; !y {
		fmt.Println("y is false")
	}
}

func Switch() {
	// switch文で変数宣言するパターン
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	//if-elseの代わりにswitch文を使うケース
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evning.")
	}
}

func Defer() {
	fmt.Println("count down!")

	for i := 0; i < 10; i++ {
		// defer は最後に実行される
		// defer の中身は即座に評価される
		// defer の中身はスタックされる
		defer fmt.Println(i)
	}

	fmt.Println("get start!")
}

func MoreTypes() {
	// ポインタ ///////////////////////////////////////////////////////////////////
	i := 1
	pointer := &i     // & でiのポンタ(アドレスと型情報)を取得
	value := *pointer // * でポインタの指す先の値を取得
	var j *int        // *+型で、この変数は型のポインタであることを宣言する。↑の値を示す*とは別ものだと思ったほうがよい。
	fmt.Println(value, *j)

	// 構造体 /////////////////////////////////////////////////////////////////////
	type Vertex struct {
		X int
		Y int
	}
	// 構造体の基本的な使い方
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 構造体をポインタで使用する方法
	vp := Vertex{1, 2}
	p := &vp
	p.X = 1e9 // (*p).X とも書けるが、冗長なのでこういう書き方ができるようになっている。
	fmt.Println(vp)

	// 構造体の初期化
	v2 := Vertex{X: 1}
	fmt.Println(v2) // {1 0}

	// array //////////////////////////////////////////////////////////////////////
	// array の基本パターン
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
	// array の宣言と初期化を同時に行うパターン
	b := [6]int{1, 4, 9, 16, 25, 36}
	fmt.Println(b)

	// slices /////////////////////////////////////////////////////////////////////
	// slice の宣言
	var s []int = b[1:4]
	fmt.Println(s) // [4, 9, 16]
	// 注意!: slice は array への参照のようなもの。slice の要素を変更すると array の要素が変更される。

	// 範囲指定の表現パターン
	s = s[:2]
	s = s[1:]
	s = s[:]

	// 注意: array & slice
	arr := [3]int{1, 2, 3} // これは array. 長さの指定がある。
	sli := []int{4, 5, 6}  // これは slice. 長さの指定がない。 4,5,6のarray を作り、それを参照するslice を作成する。
	fmt.Println(arr, sli)

	// 動的配列の作成
	sli2 := make([]int, 0, 5) // 5要素のarrayを参照する長さ0のslice
	sli3 := make([]int, 5)    // 5要素のarrayを参照する長さ5のslice
	fmt.Println(sli2, sli3)

	// slice に要素を追加する
	sli4 := make([]int, 1)
	sli4 = append(sli4, 1, 2, 3) // sli4 == [0, 1, 2, 3]

	// for range: slice や map でイテレートするために使う
	for i, v := range sli {
		fmt.Printf("index=%d value=%d", i, v)
	}

	// map ////////////////////////////////////////////////////////////////////////////
	type Geo struct {
		Latitude, Longitude float64
	}

	// make で生成するパターン
	var m map[string]Geo
	m = make(map[string]Geo)
	m["Bell Labs"] = Geo{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	// 宣言時に初期化するパターン
	var m2 = map[string]Geo{
		"Bell Labls": Geo{40.68433, -74.39967},
		"Google":     Geo{37.42202, -122.08408},
	}
	fmt.Println(m2)

	// 宣言時に初期化するパターン(構造体の方を推定させるパターン)
	var m3 = map[string]Geo{
		"Bell Labls": {40.68433, -74.39967},
		"Google":     {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// 要素の追加
	m4 := make(map[string]int)
	m4["test"] = 1
	// 要素の更新
	m4["test"] = 2
	// 要素の削除
	delete(m4, "test")
	// 要素の取得
	value, ok := m4["test"]
	fmt.Println("the value: ", value, "present?", ok)

}

// funcも変数 ///////////////////////////////////////////////////////////////////////////
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func fun() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))
}

// methods /////////////////////////////////////////////////////////////////////////////
type Vertex struct {
	X, Y float64
}

// 変数レシーバのmethod。
// method は、同じパッケージで定義された型についてのみ定義できる。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバのmethod。
// 型のインスタンスの中の値を更新するときに使う。
// こちらのほうが一般的。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 注意: 一般的には、値レシーバまたはポインタレシーバのどちらかですべてのメソッドを与え、
//       混在させるべきではない。
func Methods() {
	v := Vertex{3, 4}
	v.Scale(10) // (&v).Scale(10) と書いてもOK
	v.Abs()
}
