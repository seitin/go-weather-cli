package utils

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "io"
  "go-weather/src/types"
)

const IPAPIUrl string = "http://ip-api.com/json/" 


func GetLocationInfo() types.LocationInfo {
  var ipJson types.LocationInfo

  res, err := http.Get(IPAPIUrl)
  if err != nil {
    fmt.Printf("error while getting external IP")
    log.Fatal(err)
  }

  bytebody, err := io.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  json.Unmarshal(bytebody, &ipJson)
  return ipJson
}


