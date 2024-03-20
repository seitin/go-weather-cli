package utils

import (
	"encoding/json"
	t "go-weather/src/types"
	"io"
	"log"
	"net/http"
)

const WEATHER_URL string = "https://api.weatherapi.com/v1/current.json"

type GetWeatherParams struct {
	IPAddress string
	APIKey    string
	Lang      string
}

func GetWeatherInfo(params GetWeatherParams) t.WeatherInfo {
	var q string = params.IPAddress
	var queryString string
	var result t.WeatherInfo

	queryString = "?q=" + q + "&key=" + params.APIKey + "&lang=" + params.Lang

	var url string = WEATHER_URL + queryString
	log.Print(url)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	bytebody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytebody, &result)
	if err != nil {
		log.Fatal(err)
	}
	CreateLastInfoFile(result)
	return result
}
