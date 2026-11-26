package services

import (
	"errors"
	"time"

	"meteo/api"
	"meteo/models"
)

func GetDays(commune models.Commune) ([]models.Day, error) {
	if len(commune.Centre.Coordinates) != 2 {
		return nil, errors.New("coordonnées manquantes pour " + commune.Code)
	}
	lon := commune.Centre.Coordinates[0]
	lat := commune.Centre.Coordinates[1]

	forecast, err := api.GetForecast(lat, lon)
	if err != nil {
		return nil, err
	}
	return BuildDays(forecast.Daily)
}

func BuildDays(daily models.Daily) ([]models.Day, error) {
	n := len(daily.Time)
	if len(daily.TemperatureMin) != n || len(daily.TemperatureMax) != n || len(daily.WeatherCode) != n {
		return nil, errors.New("les tableaux de prévisions n'ont pas la même longueur")
	}

	var days []models.Day
	for i := 0; i < n; i++ {
		date, err := time.Parse("2006-01-02", daily.Time[i])
		if err != nil {
			return nil, err
		}
		day := models.Day{
			Date: date,
			Min:  daily.TemperatureMin[i],
			Max:  daily.TemperatureMax[i],
			Code: daily.WeatherCode[i],
		}
		days = append(days, day)
	}
	return days, nil
}
