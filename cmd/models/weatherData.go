package models

import (
	"encoding/json"
	"math"
)

type Temperature int

func (t *Temperature) UnmarshalJSON(data []byte) error {
	var temp float64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*t = Temperature(math.Round(temp))
	return nil
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius  Temperature `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	HasData bool 
}

