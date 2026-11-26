package routes

import (
	"net/http"

	"meteo/controllers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", controllers.Home)
	mux.HandleFunc("GET /ville", controllers.Ville)
	mux.HandleFunc("GET /meteo", controllers.Meteo)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}
