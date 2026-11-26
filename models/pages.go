package models

type HomePage struct {
	Query string
	Error string
}

type VillesPage struct {
	Query    string
	Communes []Commune
}

type MeteoPage struct {
	Commune Commune
	Days    []Day
	MapURL  string
}
