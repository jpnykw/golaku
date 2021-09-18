package main
import "golang.org/x/tour/reader"

type MyReader struct{}

func (_ MyReader) Read(buffer []byte) (int, error) {
  for i := range buffer {
    buffer[i] = 'A'
  }
  return len(buffer), nil
}

func main() {
  reader.Validate(MyReader{})
}

