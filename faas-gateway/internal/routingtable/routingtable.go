package routingtable

import (
  "context"
  "time"
  "fmt"
  "strconv"
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
  lname, lport := "faas.name", "faas.port"
  filters := filters.NewArgs()
  filters.Contains(lname)
  filters.Contains(lport)
  for {
    containers, err := rt.Client.ContainerList(context.Background(), types.ContainerListOptions{
      Filters: filters,
    })
    if err != nil {
      panic(err)
    }
    services := make(map[string]service.Service)
    for _, container := range containers {
      labels := container.Labels
      name, state := labels[lname], container.State
      port, _ := strconv.Atoi(labels[lport])
      network := container.NetworkSettings
      ipAddress := network.Networks["bridge"].IPAddress
      srv, exists := services[name]
      if !exists {
        srv = service.Service {
          Name: name,
          Total: 0,
          Available: 0,
          Port: port,
          IPAddresses: make([]string, 0),
        }
      }
      srv.Total += 1
      if state == "running" {
        srv.Available += 1
        srv.IPAddresses = append(srv.IPAddresses, ipAddress)
      }
      services[name] = srv
      message := fmt.Sprintf("%s", srv) 
      rt.Messages <- message
    }
    rt.Services = services
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
