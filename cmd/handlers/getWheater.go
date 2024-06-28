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

func GetWheater(c echo.Context)  error{
	var weatherData models.WeatherData
	
	if err := godotenv.Load(); err != nil{
		return c.JSON(http.StatusInternalServerError, "Erro to load .Env files")
	}
	apiKey := os.Getenv("API_KEY")
	
	city := c.Param("city")

	spacedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&lang=pt_br&appid=%s&units=metric", spacedCity, apiKey)

	data, err := http.Get(url)
	if err != nil{
		return  err
	}
	defer data.Body.Close()
	
	if err := json.NewDecoder(data.Body).Decode(&weatherData); err != nil{
		return  err
	}

	return c.JSON(http.StatusOK, weatherData)
}