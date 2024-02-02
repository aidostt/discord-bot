package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

// WeatherData holds the structured response from the weather API.
type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

// WeatherClient encapsulates the functionality to fetch weather data.
type WeatherClient struct {
	APIKey string
}

// NewWeatherClient creates a new instance of WeatherClient.
func NewWeatherClient() *WeatherClient {
	return &WeatherClient{
		APIKey: os.Getenv("OPENWEATHERMAP_API_KEY"), // Ensure you have set this environment variable
	}
}

// GetWeather fetches the weather for a given city.
func (client *WeatherClient) GetWeather(city string) (*WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, client.APIKey)

	// Create a channel to receive the API response
	ch := make(chan *WeatherData, 1)
	chErr := make(chan error, 1)

	// Make the API request in a goroutine
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			chErr <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			chErr <- errors.New("failed to find the city")
			return
		}

		var data WeatherData
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			chErr <- err
			return
		}

		ch <- &data
	}()

	// Use a select statement to implement the timeout
	select {
	case data := <-ch:
		return data, nil
	case err := <-chErr:
		return nil, err
	case <-time.After(5 * time.Second):
		// Return an error if the API call doesn't complete within 5 seconds
		return nil, errors.New("request timed out or city not found")
	}
}
