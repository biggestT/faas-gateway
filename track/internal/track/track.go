package track

import (
  "github.com/biggestT/tracker/internal/location"
  "encoding/csv"
  "bufio"
  "io"
  "os"
)

type TrackMap = map[string]*Track
type headerMap = map[string]int

type Track struct {
  Location *location.Location `json:"location"`
  Name string `json:"name"`
}

func parseRecord(r []string, h headerMap) (*Track, string) {
  t := Track{
    location.New(
      r[h["latitidue"]],
      r[h["longitude"]],
    ),
    r[h["name"]],
  }
  return &t, r[h["code"]]
}

func New(file string) TrackMap {
  f, _ := os.Open(file)
  tmap := make(TrackMap)
  var header headerMap 
  csvReader := csv.NewReader(bufio.NewReader(f))
  for {
    record, err := csvReader.Read()
    if err == io.EOF {
      break
    }
    if header == nil {
      header = make(headerMap)
      for i, c := range record {
        header[c] = i
      }
    } else {
      track, code := parseRecord(record, header)
      tmap[code] = track
    }
  }
  return tmap
}
