package service

import (
  "strings"
)

type Service struct {
  Name string
  Total int
  Available int
  Port string
  IPAddresses []string
}

func GetHost(service *Service) string {
  parts := []string{service.IPAddresses[0], service.Port}
  return strings.Join(parts, ":")
}
