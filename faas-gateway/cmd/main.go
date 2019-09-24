package main

import (
  "github.com/biggestT/faas-gateway/internal/routingtable"
  "github.com/biggestT/faas-gateway/internal/proxy"
  "net/http"
	"fmt"
)

func log(rt *routingtable.RoutingTable) {
  for {
    msg := <- rt.Messages
    fmt.Println(msg)
  }
}

func main(){
  fmt.Println("Gateway initiating")
  rt, _ := routingtable.NewRoutingTable()
  fmt.Println("Service discovery initiated")
  go log(rt)
  http.Handle("/", proxy.ProxyServer(rt))
  fmt.Println("Proxy server started")
  if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
