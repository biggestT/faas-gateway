// continuously keeps track of available services
package routingtable

import (
  "container/list"
  "context"
  "time"
  "fmt"
  "github.com/biggestT/faas-gateway/internal/service"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/filters"
  "github.com/docker/docker/client"
)

type serviceMap = map[string]*service.Service

type RoutingTable struct {
  Routes serviceMap 
  Messages chan string
  dockerClient *client.Client
}

func UpdateMessages(a serviceMap, b serviceMap) *list.List {
  msgs := list.New() 
  for name, _ := range b {
    if a[name] == nil {
      msgs.PushFront(fmt.Sprintf("service discovered: %s", name))
    }
  }
  for name, _ := range a {
    if b[name] == nil {
      msgs.PushFront(fmt.Sprintf("service removed: %s", name))
    }
  }
  return msgs
}

func (r *RoutingTable) poll(freq time.Duration) {
  lname, lport := "faas.name", "faas.port"
  filter := filters.NewArgs()
  filter.Add("label", "faas.app=true")
  for {
    containers, err := r.dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
      Filters: filter,
    })
    if err != nil {
      panic(err)
    }
    routes := make(map[string]*service.Service)
    for _, container := range containers {
      labels := container.Labels
      name, port, state := labels[lname], labels[lport], container.State
      // docker API adds a forward slash to the name
      containerName := container.Names[0][1:]
      srv, exists := routes["/" + name]
      if !exists {
        srv = service.New(port)
      }
      if state == "running" {
        srv.AddHost(containerName)
      }
      routes["/" + name] = srv
    }
    msgs := UpdateMessages(r.Routes, routes)
    for e := msgs.Front(); e != nil; e = e.Next() {
      r.Messages <- e.Value.(string)
    }
    r.Routes = routes
    time.Sleep(time.Second * freq)
  }
}

func NewRoutingTable(freq int) (*RoutingTable, error) {
  rt := new(RoutingTable)
  cli, err := client.NewClientWithOpts(client.FromEnv)
  rt.dockerClient = cli
  rt.Routes = make(map[string]*service.Service)
  rt.Messages = make(chan string)
  go rt.poll(time.Duration(freq))
  return rt, err
}
