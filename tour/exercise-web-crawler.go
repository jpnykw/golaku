package main
import (
  "fmt"
  "sync"
)

type Fetcher interface {
  Fetch(url string) (body string, urls []string, err error)
}

type FetchedLinks struct {
  mutex sync.Mutex
  fetched map[string]bool
}

var fetchedUrls FetchedLinks = FetchedLinks {
  fetched: make(map[string]bool),
}

func Crawl(url string, depth int, fetcher Fetcher) {
  if depth <= 0 {
    return
  }

  if (func () bool {
    fetchedUrls.mutex.Lock()
    defer fetchedUrls.mutex.Unlock()

    if _, ok := fetchedUrls.fetched[url]; !ok {
      fetchedUrls.fetched[url] = true
      return false
    } else {
      return true
    }
  })() {
    return
  }

  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    Crawl(u, depth-1, fetcher)
  }

  return
}

func main() {
  Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
    return "", nil, fmt.Errorf("not found: %s", url)
  }

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
  "https://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "https://golang.org/pkg/",
      "https://golang.org/cmd/",
    },
  },
  "https://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "https://golang.org/",
      "https://golang.org/cmd/",
      "https://golang.org/pkg/fmt/",
      "https://golang.org/pkg/os/",
    },
  },
  "https://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "https://golang.org/",
      "https://golang.org/pkg/",
    },
  },
  "https://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "https://golang.org/",
      "https://golang.org/pkg/",
    },
  },
}

