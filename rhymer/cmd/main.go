package main

import (
  "github.com/biggestT/rhymer/internal/dict"
  "github.com/biggestT/rhymer/internal/api"
  "net/http"
  "fmt"
)

func main(){
  fmt.Println("rhymer initiating")
  dict := dict.New("data/wordlist.txt")

  http.Handle("/", api.ApiHandler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
