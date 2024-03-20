package utils

import (
	wt "go-weather/src/types"
	"io/ioutil"
	"os"
	"testing"
)
func TestGenerateTempFilePath(t *testing.T) {
	expected := os.TempDir() + TEMP_FILE_NAME
	got := GenerateTempFilePath()

	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestCreateLastInfoFile(t *testing.T) {
	weather := wt.WeatherInfo{
		Location: wt.Location{
			Name: "San Francisco",
		},
	}

	CreateLastInfoFile(weather)

	filepath := GenerateTempFilePath()
	_, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Errorf("Failed to create file: %s", err)
	}

	// Clean up after test
	os.Remove(filepath)
}

func TestGetCachedWeatherInfo(t *testing.T) {
	weather := wt.WeatherInfo{
		Location: wt.Location{
			Name: "San Francisco",
		},
	}
	CreateLastInfoFile(weather)

	got := GetCachedWeatherInfo()
	if got != weather {
		t.Errorf("Expected %v, got %v", weather, got)
	}

	// Clean up after test
	os.Remove(GenerateTempFilePath())
}
