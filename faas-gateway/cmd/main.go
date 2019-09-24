package main

import (
  "github.com/biggestT/faas-gateway/internal/routingtable"
	"fmt"
)

func main(){
  fmt.Println("Gateway Started")
  rt, _ := routingtable.NewRoutingTable()
  for {
    msg := <- rt.Messages
    fmt.Println(msg)
  }
  select {}
}
