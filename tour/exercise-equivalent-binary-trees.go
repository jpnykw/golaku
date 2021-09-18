package main
import (
  "golang.org/x/tour/tree"
  "fmt"
)

func Walk(t *tree.Tree, ch chan int) {
  if t != nil {
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
  }
}

func WalkAndClose(t *tree.Tree, ch chan int) {
  Walk(t, ch)
  close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
  ch_for_t1, ch_for_t2 := make(chan int), make(chan int)

  // I wrapped Walk function cause I didn't get of a good way to close the channel
  go WalkAndClose(t1, ch_for_t1)
  // close(ch_for_t1)
  go WalkAndClose(t2, ch_for_t2)
  // close(ch_for_t2)

  for {
    v1, is_open_ch1 := <- ch_for_t1
    v2, is_open_ch2 := <- ch_for_t2

    switch {
      case !is_open_ch1, !is_open_ch2:
        return is_open_ch1 == is_open_ch2
      case v1 != v2:
        return false
    }
  }
}

func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}

