package main
import (
	"golang.org/x/tour/wc"
	"strings"
)

type Res = map[string]int

func WordCount(s string) (words Res) {
	words = make(Res)
	for _, word := range strings.Fields(s) {
		words[word] += 1
	}
	return
}

func main() {
	wc.Test(WordCount)
}

