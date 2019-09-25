// passes incoming requests to appropiate service if possible
package proxy

import (
  "net/http"
  "net/http/httputil"
  "github.com/biggestT/faas-gateway/internal/routingtable"
)

type proxyHandler struct {
  routing *routingtable.RoutingTable
  corsOrigin string
}

func (f *proxyHandler) enableCors(w *http.ResponseWriter) {
  header := (*w).Header()
	header.Set("Access-Control-Allow-Origin", f.corsOrigin)
	header.Set("Access-Control-Allow-Headers", "*")
  header.Set("Access-Control-Allow-Methods", "GET, POST, HEAD")
}

func (f *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  srv, exists := f.routing.Routes[r.URL.Path]
  if !exists {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  srvHost := srv.NextHost()
  r.URL.Scheme = "http"
  r.URL.Host = srvHost
  r.URL.Path = ""
  f.enableCors(&w)
  proxy := httputil.NewSingleHostReverseProxy(r.URL)
  proxy.ServeHTTP(w, r)
}

func ProxyServer(
  routing *routingtable.RoutingTable,
  corsOrigin string,
) http.Handler {
  return &proxyHandler{routing, corsOrigin}
}


