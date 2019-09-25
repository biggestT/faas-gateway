// returns appropriate rhymes for requested name
package api

import (
  "net/http"
  "github.com/biggestT/rhymer/internal/dict"
)

type requestHandler struct {
  dict *dict.Dict
}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  names, ok := r.URL.Query()["name"]
  if !ok || len(names) < 1 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  rhyme := h.dict.NextRhyme(names[0])
  w.Write([]byte(rhyme))
}

func Api(
  dict *dict.Dict,
) http.Handler {
  return &requestHandler{dict}
}
