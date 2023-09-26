package types

import (

)

type LocationInfo struct {
  Lat float32 `json:"lat"`
  Lon float32 `json:"lon"`
  Query string `json:"query"`
}

type WeatherInfo struct {
  Location struct {
    Name string `json:"name"`
    Region string
    Country string
  }
  Current struct {
    TempC float32 `json:"temp_c"`
    TempF float32 `json:"temp_f"`
    Condition struct {
      Text string
    }
  }
}
