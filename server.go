package main

import (
  "flag"
  "fmt"
  "strconv"
  "net/http"
  "os/exec"
)

const default_port = "8080"

func handler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hey guys, what's up?"))
}

func main() {
  arg := flag.String("port", default_port, "port to listen on localhost")
  flag.Parse()
  port, _ := strconv.Atoi(*arg)
  
  e := exec.Command("open", fmt.Sprintf("http://0.0.0.0:%d", port)).Start()
  if e != nil {
    panic(e)
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

