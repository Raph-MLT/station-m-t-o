package helpers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	funcs := template.FuncMap{
		"frDate":       FrDate,
		"temp":         Temp,
		"weatherLabel": WeatherLabel,
		"weatherIcon":  WeatherIcon,
	}
	templates = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))
}

func Render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}
