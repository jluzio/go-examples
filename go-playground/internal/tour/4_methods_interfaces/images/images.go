package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func main() {
	// main_images()
	exercise_images()
}

func main_images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct {
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0.0, 0.0},
		image.Point{200.0, 100.0},
	}
}
func (img Image) At(x, y int) color.Color {
	var r uint8 = uint8((x + y) / 2)
	var g uint8 = uint8(x * y)
	var b uint8 = uint8(x ^ y)
	var a uint8 = 255
	return color.RGBA{r, g, b, a}
}

// check in browser with "data:image/png;base64," + <result without IMAGE:>
func exercise_images() {
	m := Image{}
	pic.ShowImage(m)
}
