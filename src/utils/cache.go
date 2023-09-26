package utils

import (
	"encoding/json"
	t "go-weather/src/types"
	"log"
	"os"
)

const TEMP_FILE_NAME string = "go-weather.tmp"

func GenerateTempFilePath() string {
  var filepath string = os.TempDir() + TEMP_FILE_NAME
  return filepath
}

func CreateLastInfoFile(weather t.WeatherInfo) {
  bin, err := json.Marshal(weather)
  if err != nil {
    log.Fatal(err)
  }
  var filepath string = GenerateTempFilePath()
  err = os.WriteFile(filepath, bin, 0777)

  if err != nil {
    log.Fatal(err)
  }
}

func GetCachedWeatherInfo() t.WeatherInfo {
  var filepath string = GenerateTempFilePath()
  var weather t.WeatherInfo
  bin, err := os.ReadFile(filepath)
  if err != nil {
    return t.WeatherInfo{}
  }
  
  err = json.Unmarshal(bin, &weather)
  if err != nil {
    return t.WeatherInfo{}
  }

  return weather
}
