package main

import (
  "github.com/biggestT/faas-gateway/internal/routingtable"
  "github.com/biggestT/faas-gateway/internal/proxy"
  "net/http"
  "os"
  "fmt"
)

func log(rt *routingtable.RoutingTable) {
  for {
    msg := <- rt.Messages
    fmt.Println(msg)
  }
}

func main(){
  fmt.Println("gateway initiating")
  rt, _ := routingtable.NewRoutingTable()
  fmt.Println("service discovery initiated")
  go log(rt)
  corsOrigin := os.Getenv("CORS_ORIGIN")
  http.Handle("/", proxy.ProxyServer(rt, corsOrigin))
  fmt.Println("proxy server started")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
