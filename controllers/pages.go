package controllers

import (
	"log"
	"net/http"

	"meteo/api"
	"meteo/helpers"
	"meteo/models"
	"meteo/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, http.StatusOK, "index.html", models.HomePage{})
}

func Ville(w http.ResponseWriter, r *http.Request) {
	q := services.CleanCity(r.URL.Query().Get("q"))

	message := services.ValidateCity(q)
	if message != "" {
		helpers.Render(w, http.StatusBadRequest, "index.html", models.HomePage{Query: q, Error: message})
		return
	}

	communes, err := api.SearchCommunes(q)
	if err != nil {
		log.Println(err)
		helpers.Render(w, http.StatusServiceUnavailable, "error.html", "Service momentanément indisponible.")
		return
	}

	if len(communes) == 0 {
		helpers.Render(w, http.StatusNotFound, "error.html", "Aucune commune ne correspond à « "+q+" ».")
		return
	}

	helpers.Render(w, http.StatusOK, "villes.html", models.VillesPage{Query: q, Communes: communes})
}

func Meteo(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		helpers.Render(w, http.StatusNotFound, "error.html", "Cette commune n'existe pas.")
		return
	}

	commune, err := api.GetCommune(code)
	if err == api.ErrNotFound {
		helpers.Render(w, http.StatusNotFound, "error.html", "Cette commune n'existe pas.")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.Render(w, http.StatusServiceUnavailable, "error.html", "Service momentanément indisponible.")
		return
	}

	days, err := services.GetDays(commune)
	if err != nil {
		log.Println(err)
		helpers.Render(w, http.StatusServiceUnavailable, "error.html", "Service momentanément indisponible.")
		return
	}

	helpers.Render(w, http.StatusOK, "meteo.html", models.MeteoPage{Commune: commune, Days: days})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, http.StatusNotFound, "error.html", "Cette page n'existe pas.")
}
