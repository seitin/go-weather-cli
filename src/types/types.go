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
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
    TempF            float32 `json:"temp_f"`
		TempC            float32 `json:"temp_c"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
      Code int    `json:"code"`
    } `json:"condition"`
    WindMph    float64 `json:"wind_mph"`
    WindKph    float32     `json:"wind_kph"`
    WindDegree float32     `json:"wind_degree"`
    WindDir    string  `json:"wind_dir"`
    PressureMb float32     `json:"pressure_mb"`
    PressureIn float32 `json:"pressure_in"`
    PrecipMm   float32     `json:"precip_mm"`
    PrecipIn   float32     `json:"precip_in"`
    Humidity   float32     `json:"humidity"`
    Cloud      float32     `json:"cloud"`
    FeelslikeC float32 `json:"feelslike_c"`
    FeelslikeF float64 `json:"feelslike_f"`
    VisKm      float32     `json:"vis_km"`
    VisMiles   float32     `json:"vis_miles"`
    Uv         float32     `json:"uv"`
    GustMph    float32     `json:"gust_mph"`
    GustKph    float64 `json:"gust_kph"`
  } `json:"current"`
}
