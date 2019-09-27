package main

import (
  "github.com/biggestT/rhymer/internal/dict"
  "github.com/biggestT/rhymer/internal/api"
  "net/http"
  "log"
  "os"
  "fmt"
)
func main(){
  fmt.Println("rhymer initiating")
  dict := dict.New("data/wordlist.txt")
  logger := log.New(os.Stdout, "http: ", log.LstdFlags)
  apiHandler := api.Api(dict)
  http.Handle("/", logging(logger)(apiHandler))
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

// from https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7
func logging(logger *log.Logger) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      defer func() {
        logger.Println(r.Method, r.URL.Path, r.URL.Query())
      }()
      next.ServeHTTP(w, r)
    })
  }
}
