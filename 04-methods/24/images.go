// なんとなくわかりました。
package main

import (
	"fmt"
	"image"
)

func main() {
	// func Rect(x0, y0, x1, y1 int) Reactangle : 長方形 {Pt(x0, y0), Pt(x1, y1)} の省略形。
	// type Reactangle : 四角形。その境界が自分自身である Image。
	// func NewRGBA(r Rectangle) *RGBA : 指定された範囲の新しいRGBA画像を返す。
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// func Bounds() : 画像の領域を取得する。
	fmt.Println(m.Bounds())

	// func At(x, y int) Color : x,y で指定した位置(ピクセル)の色を取得する。
	// func (c *Uniform) RGBA() (r,g,b,a uint32) : RGBAで色を表す数値を取得する。
	// type Uniform struct {C color.Color} : 画像。color.Color, color.Model,および Image の各インターフェースを実装する。
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.At(101, 101).RGBA())

}
