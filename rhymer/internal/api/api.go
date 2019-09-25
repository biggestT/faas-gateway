// passes incoming requests to appropiate service if possible
package proxy

import (
  "net/http"
  "net/http/httputil"
  "github.com/biggestT/rhymer/internal/dict"
)

type apiHandler struct {
  rhymer *dict.RoutingTable
}

func (f *proxyHandler) ApiHandler(w http.ResponseWriter, r *http.Request) {
  srv, exists := f.routing.Routes[r.URL.Path]
  if !exists {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  srvHost := srv.NextHost()
  r.URL.Scheme = "http"
  r.URL.Host = srvHost
  r.URL.Path = ""
  proxy := httputil.NewSingleHostReverseProxy(r.URL)
  proxy.ServeHTTP(w, r)
}

func ApiHandler(
  routing *routingtable.RoutingTable,
) http.Handler {
  return &proxyHandler{routing}
}
