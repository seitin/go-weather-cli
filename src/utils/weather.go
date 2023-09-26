package utils

import (
	"encoding/json"
	// "fmt"
	t "go-weather/src/types"
	"io"
	"log"
	"net/http"
)

const WEATHER_URL string = "https://api.weatherapi.com/v1/current.json" 

type GetWeatherParams struct {
  IPAddress string
  APIKey string
  Lang string
}

func GetWeatherInfo(params GetWeatherParams) t.WeatherInfo {
  var q string = params.IPAddress
  // var location t.LocationInfo = params.Location
  var queryString string 
  var result t.WeatherInfo
  
  // if location.Lat != 0 && location.Lon != 0 {
  //   q = fmt.Sprintf("%f", location.Lat) + " " + fmt.Sprintf("%f", location.Lon)
  // }

  queryString = "?q=" + q + "&key=" + params.APIKey + "&lang=" + params.Lang
  
  var url string = WEATHER_URL + queryString
  // fmt.Printf("", url)
  res, err := http.Get(url)

  if err != nil {
    log.Fatal(err)
  }

  bytebody, err := io.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  json.Unmarshal(bytebody, &result)
  return result
}


