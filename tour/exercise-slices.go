package main
import "golang.org/x/tour/pic"

func Pic(dx, dy int) (img [][]uint8) {
  img = make([][]uint8, dy)
  for x := 0; x < dy; x++ {
    img[x] = make([]uint8, dx)
  }

  for y := 0; y < dy; y++ {
    for x := 0; x < dx; x++ {
      // img[y][x] = uint8(float64((x + y) / 2)) // (x + y) / 2
      // img[y][x] *= uint8(x * y) // x * y
      img[y][x] = uint8(x ^ y) // x ^ y
    }
  }

  return
}

func main() {
  pic.Show(Pic)
}

