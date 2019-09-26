// returns neat SVG:s for requested track
package api

import (
  "net/http"
  "encoding/json"
  "github.com/biggestT/tracker/internal/track"
)

type requestHandler struct {
  trackMap track.TrackMap
}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  code, ok := r.URL.Query()["code"]
  if !ok || len(code) < 1 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  track, exists := h.trackMap[code[0]]
  if !exists {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  w.Header().Add("Content-Type", "application/json")
  resp, _ := json.Marshal(track)
  w.Write([]byte(string(resp)))
}

func Api(
  trackMap track.TrackMap,
) http.Handler {
  return &requestHandler{trackMap}
}
