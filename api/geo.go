package api

import (
	"net/url"

	"meteo/models"
)

const geoFields = "code,nom,codesPostaux,population,centre"

func SearchCommunes(nom string) ([]models.Commune, error) {
	params := url.Values{}
	params.Set("nom", nom)
	params.Set("fields", geoFields)
	params.Set("boost", "population")
	params.Set("limit", "10")

	var communes []models.Commune
	err := getJSON("https://geo.api.gouv.fr/communes?"+params.Encode(), &communes)
	return communes, err
}

func GetCommune(code string) (models.Commune, error) {
	params := url.Values{}
	params.Set("fields", geoFields)

	var commune models.Commune
	err := getJSON("https://geo.api.gouv.fr/communes/"+url.PathEscape(code)+"?"+params.Encode(), &commune)
	return commune, err
}
