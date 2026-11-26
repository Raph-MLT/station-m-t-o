package models

import "time"

type Forecast struct {
	Daily Daily `json:"daily"`
}

type Daily struct {
	Time           []string  `json:"time"`
	TemperatureMax []float64 `json:"temperature_2m_max"`
	TemperatureMin []float64 `json:"temperature_2m_min"`
	WeatherCode    []int     `json:"weather_code"`
}

type Day struct {
	Date     time.Time
	Min      float64
	Max      float64
	Code     int
	BarLeft  float64
	BarWidth float64
}
