package main

import (
  "image"
  "image/color"
  "golang.org/x/tour/pic"
)

type Image struct{}

func (_ Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, 128, 128)
}

func (_ Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (_ Image) At(x, y int) color.Color {
  // return color.RGBA{uint8(x), uint8(y), 255, 255} // last picture generator corresponds
  return color.RGBA{uint8((y * (x ^ y)) / 16), uint8(y * 2), uint8(y | (x << 1)), 255}
}

func main() {
  m := Image{}
  pic.ShowImage(m)
}

