package utils

import (
	"encoding/json"
	t "go-weather/src/types"
	"log"
	"os"
)

const TEMP_FILE_NAME string = "go-weather.tmp"

// GenerateTempFilePath is a function that generates a temporary file path.
// It concatenates the system's temporary directory path with a predefined temporary file name.
// The function does not take any parameters.
// Returns:
//   string: The full path to the temporary file.
func GenerateTempFilePath() string {
	var filepath string = os.TempDir() + TEMP_FILE_NAME
	return filepath
}
// CreateLastInfoFile is a function that creates a file with the last fetched weather information.
// It takes a WeatherInfo struct as parameter, serializes it to JSON and writes it to a temporary file.
// If the serialization or the file writing fails, the function will log the error and stop the execution.
// Parameters:
//   weather (t.WeatherInfo): The weather information to be cached.
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

// GetCachedWeatherInfo is a function that retrieves the last fetched weather information from a temporary file.
// It reads the file, deserializes the JSON content to a WeatherInfo struct and returns it.
// If the file reading or the deserialization fails, the function will return an empty WeatherInfo struct.
// Returns:
//   t.WeatherInfo: The cached weather information.
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
