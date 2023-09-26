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
	"time"
	externalip "github.com/glendc/go-external-ip"
)

const CACHE_LIMIT_IN_MINUTES = 15
type GoWeatherSettings struct {
  WeatherAPIKey string
  Language string
}

func ReadSettings() GoWeatherSettings {
  var settings GoWeatherSettings
  userHomeDir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  }
  bytearray, err := os.ReadFile(userHomeDir + "/.goweather/conf.json")
  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(bytearray, &settings)
  if err != nil {
    log.Print("Couldn't find the configuration file that should be located in $HOME/.goweather/conf.json")
    log.Fatal(err)
  }
  return settings
}


func IsCacheInvalid(cachedWeatherInfo types.WeatherInfo) bool {
  var unixmilliNow int64 = time.Now().UTC().UnixMilli()
  var lastupdate int64 = int64(cachedWeatherInfo.Current.LastUpdatedEpoch) * 1000
  var threshold int64 = lastupdate + 1000 * 60 * CACHE_LIMIT_IN_MINUTES
  var isCacheInvalid bool = unixmilliNow > threshold || (lastupdate == 0)
  return isCacheInvalid
}

func PrintWeather(weather types.WeatherInfo) {
  fmt.Printf("%s | %s | %g°C | %s", weather.Location.Name, weather.Current.Condition.Text, weather.Current.TempC, weather.Location.Localtime)
}

func main() {
  var cachedWeatherInfo types.WeatherInfo = utils.GetCachedWeatherInfo()
  if !IsCacheInvalid(cachedWeatherInfo) {
    PrintWeather(cachedWeatherInfo)
    return
  }


  var settings GoWeatherSettings = ReadSettings()
  decodedApiKey, err := base64.StdEncoding.DecodeString(settings.WeatherAPIKey)
  if err != nil {
    log.Fatal(err)
  }

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
  PrintWeather(weather)
}
