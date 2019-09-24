package proxy

import (
  "fmt"
  "net/http"
  "net/http/httputil"
  "github.com/biggestT/faas-gateway/internal/routingtable"
  "github.com/biggestT/faas-gateway/internal/service"
)

type proxyHandler struct {
  routing *routingtable.RoutingTable
}

func (f *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  srv, exists := f.routing.Routes[r.URL.Path]
  if !exists {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  fmt.Println(r.URL.Path)
  srvHost := service.GetHost(&srv)
  r.URL.Scheme = "http"
	r.URL.Host = srvHost
	r.URL.Path = ""
	proxy := httputil.NewSingleHostReverseProxy(r.URL)
	proxy.ServeHTTP(w, r)
}

func ProxyServer(
  routing *routingtable.RoutingTable,
) http.Handler {
  return &proxyHandler{routing}
}
