package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"wheater/cmd/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetWheater(city string) (models.WeatherData, error){
	var weatherData models.WeatherData
	if err := godotenv.Load(); err != nil{
		return weatherData, err
	}
	apiKey := os.Getenv("API_KEY")
	
	spacedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&lang=pt_br&appid=%s&units=metric", spacedCity, apiKey)

	data, err := http.Get(url)
	if err != nil{
		return models.WeatherData{}, err
	}
	defer data.Body.Close()
	
	if err := json.NewDecoder(data.Body).Decode(&weatherData); err != nil{
		return models.WeatherData{}, err
	}
	
	weatherData.HasData = true
	if weatherData.Name == "" {
		weatherData.HasData = false
	}
	return weatherData, nil
}

func GetCity(c echo.Context) error {
	city := c.QueryParam("city")
	weatherData, err := GetWheater(city)

	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "index.html", weatherData)
}