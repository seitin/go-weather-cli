package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-weather/src/types"
	"go-weather/src/utils"
	"log"
	"os"
	"strings"

	externalip "github.com/glendc/go-external-ip"
)

type GoWeatherSettings struct {
  WeatherAPIKey string
  Language string
}

func ReadSettings() GoWeatherSettings {
  var settings GoWeatherSettings
  userHomeDir, err := os.UserHomeDir()
  bytearray, err := os.ReadFile(userHomeDir + "/.goweather/conf.json")
  if err != nil {
    log.Fatal(err)
  }

  json.Unmarshal(bytearray, &settings)
  return settings
}
func main() {

  var settings GoWeatherSettings = ReadSettings()
  decodedApiKey, err := base64.StdEncoding.DecodeString(settings.WeatherAPIKey)
  if err != nil {
    log.Fatal(err)
  }
  // var location types.LocationInfo = utils.GetLocationInfo()

  var apiKey string = strings.Trim(string(decodedApiKey), "\n")
  consensus := externalip.DefaultConsensus(nil, nil)

  ip, err := consensus.ExternalIP()
  if err != nil {
    log.Fatal(err)
  }

  var weatherParams = utils.GetWeatherParams {
    Lang: settings.Language,
    IPAddress: ip.String(),
    APIKey: apiKey,
  }


  var weather types.WeatherInfo = utils.GetWeatherInfo(weatherParams)
  fmt.Printf("%s | %s | %g°C", weather.Location.Name, weather.Current.Condition.Text, weather.Current.TempC)
}
