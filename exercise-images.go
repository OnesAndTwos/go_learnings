package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image struct
type Image struct{}

//ColorModel color.Model
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds Rectangle
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

//At color.Color
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
