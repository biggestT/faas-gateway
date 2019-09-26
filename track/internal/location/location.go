package location

import (
  "strconv"
)

type Location struct {
  Lat float64 `json:"latitude"`
  Long float64 `json:"longitude"`
}

func New(lt string, ln string) *Location {
  lat, _ := strconv.ParseFloat(lt, 64)
  lon, _ := strconv.ParseFloat(ln, 64)
  return &Location{
    lat, lon,
  }
}
