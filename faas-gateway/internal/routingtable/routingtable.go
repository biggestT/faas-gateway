package routingtable

import (
  "context"
  "time"
  "fmt"
  "github.com/biggestT/faas-gateway/internal/service"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/filters"
  "github.com/docker/docker/client"
)

type RoutingTable struct {
  Services map[string]service.Service
  Client *client.Client
  Messages chan string
}

func poll(rt *RoutingTable) {
  // containers
  filters := filters.NewArgs()
  filters.Contains("faas.name")
  filters.Contains("faas.port")
  for {
    fmt.Println("running poll")
    containers, err := rt.Client.ContainerList(context.Background(), types.ContainerListOptions{
      Filters: filters,
    })
    if err != nil {
      panic(err)
    }
    for _, container := range containers {
      message := fmt.Sprintf("%s %s\n", container.ID[:10], container.Image)
      rt.Messages <- message
    }
    time.Sleep(time.Second * 2)
  }
}

func NewRoutingTable() (*RoutingTable, error) {
  rt := new(RoutingTable)
  cli, err := client.NewClientWithOpts(client.FromEnv)
  rt.Client = cli
  rt.Services = make(map[string]service.Service)
  rt.Messages = make(chan string)
  go poll(rt)
  return rt, err
}
