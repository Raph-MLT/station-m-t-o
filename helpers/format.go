package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var jours = []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}

var mois = []string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."}

func FrDate(t time.Time) string {
	return jours[t.Weekday()] + " " + strconv.Itoa(t.Day()) + " " + mois[t.Month()-1]
}

func Temp(value float64) string {
	text := fmt.Sprintf("%.1f", value)
	return strings.Replace(text, ".", ",", 1) + " °C"
}

func WeatherLabel(code int) string {
	switch {
	case code == 0:
		return "Ciel dégagé"
	case code == 1:
		return "Peu nuageux"
	case code == 2:
		return "Partiellement nuageux"
	case code == 3:
		return "Couvert"
	case code == 45 || code == 48:
		return "Brouillard"
	case code >= 51 && code <= 57:
		return "Bruine"
	case code >= 61 && code <= 67:
		return "Pluie"
	case code >= 71 && code <= 77:
		return "Neige"
	case code >= 80 && code <= 82:
		return "Averses"
	case code == 85 || code == 86:
		return "Averses de neige"
	case code >= 95 && code <= 99:
		return "Orage"
	}
	return "Inconnu"
}

func WeatherIcon(code int) string {
	switch {
	case code == 0:
		return "degage"
	case code >= 1 && code <= 3:
		return "nuageux"
	case code == 45 || code == 48:
		return "brouillard"
	case code >= 51 && code <= 57:
		return "bruine"
	case code >= 61 && code <= 67:
		return "pluie"
	case code >= 71 && code <= 77:
		return "neige"
	case code >= 80 && code <= 82:
		return "averses"
	case code == 85 || code == 86:
		return "averses-neige"
	case code >= 95 && code <= 99:
		return "orage"
	}
	return "inconnu"
}
