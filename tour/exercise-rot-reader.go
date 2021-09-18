package main
import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (r rot13Reader) Read(buffer []byte) (int, error) {
  bytes, err := r.r.Read(buffer)
  if err != nil {
    return bytes, io.EOF
  }

  for i, char := range buffer {
    switch {
      // a-m (or A-M)
      case 'a' <= char && char <= 'm', 'A' <= char && char <= 'M':
        buffer[i] += 13
      // n-z (or N-Z)
      case 'n' <= char && char <= 'z', 'N' <= char && char <= 'Z':
        buffer[i] -= 13
    }
  }

  return bytes, nil
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}

