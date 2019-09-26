package main

import (
  "github.com/biggestT/tracker/internal/track"
  "github.com/biggestT/tracker/internal/api"
  "net/http"
  "fmt"
  "log"
  "os"
)
func main(){
  fmt.Println("tracker initiating")
  trackMap := track.New("data/tracks.csv")
  apiHandler := api.Api(trackMap)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
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
