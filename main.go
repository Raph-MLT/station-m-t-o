package main

import (
	"log"
	"net/http"

	"meteo/helpers"
	"meteo/routes"
)

func main() {
	helpers.LoadTemplates()
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", routes.New()))
}
