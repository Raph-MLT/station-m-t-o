package services

import (
	"errors"
	"fmt"
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

	days, err := BuildDays(forecast.Daily)
	if err != nil {
		return nil, err
	}
	AddBars(days)
	return days, nil
}

func AddBars(days []models.Day) {
	if len(days) == 0 {
		return
	}

	weekMin := days[0].Min
	weekMax := days[0].Max
	for _, day := range days {
		if day.Min < weekMin {
			weekMin = day.Min
		}
		if day.Max > weekMax {
			weekMax = day.Max
		}
	}

	ecart := weekMax - weekMin
	if ecart == 0 {
		ecart = 1
	}

	for i := range days {
		days[i].BarLeft = (days[i].Min - weekMin) / ecart * 100
		days[i].BarWidth = (days[i].Max - days[i].Min) / ecart * 100
	}
}

func MapURL(commune models.Commune) string {
	d := 0.05
	lon := commune.Centre.Coordinates[0]
	lat := commune.Centre.Coordinates[1]
	return fmt.Sprintf("https://www.openstreetmap.org/export/embed.html?bbox=%f,%f,%f,%f&layer=mapnik&marker=%f,%f",
		lon-d, lat-d, lon+d, lat+d, lat, lon)
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
