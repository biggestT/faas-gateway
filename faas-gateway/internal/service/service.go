// keeps track of one function service 
package service

import (
  "container/list"
)

type Service struct {
  port string
  hosts *list.List
  currentHost *list.Element
}

func New(port string) *Service {
  s := new(Service)
  s.hosts = list.New()
  s.port = port
  return s
}

// add host to circulation 
func (s *Service) AddHost(host string) {
  s.hosts.PushFront(host)
  s.currentHost = s.hosts.Back()
}

// circulate through available hosts
func (s *Service) NextHost() string {
  host := s.currentHost
  nextHost := host.Next()
  if nextHost == nil {
    nextHost = s.hosts.Front()
  }
  s.currentHost = nextHost 
  return host.Value.(string) + ":" + s.port
}
