package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image :
type Image struct {
	width  int
	height int
}

// Bounds :
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

// ColorModel :
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At :
func (i Image) At(x, y int) color.Color {
	// v の算出方法は適当。^ はビット演算子で XOR を意味する。
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 64}
	pic.ShowImage(m)
}
