package service

import (
  "testing"
  "github.com/biggestT/faas-gateway/internal/service"
)

func TestService(t *testing.T) {
  s := service.New("80")
  s.AddHost("test1")
  s.AddHost("test2")
  if s.NextHost() != "test1:80" {
    t.Errorf("next host failed")
  }
  if s.NextHost() != "test2:80" {
    t.Errorf("next host failed")
  }
  if s.NextHost() != "test1:80" {
    t.Errorf("next host failed")
  }
} 
