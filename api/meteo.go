package api

import (
	"net/url"
	"strconv"

	"meteo/models"
)

func GetForecast(lat float64, lon float64) (models.Forecast, error) {
	params := url.Values{}
	params.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(lon, 'f', -1, 64))
	params.Set("daily", "temperature_2m_max,temperature_2m_min,weather_code")
	params.Set("timezone", "auto")
	params.Set("forecast_days", "7")

	var forecast models.Forecast
	err := getJSON("https://api.open-meteo.com/v1/forecast?"+params.Encode(), &forecast)
	return forecast, err
}
