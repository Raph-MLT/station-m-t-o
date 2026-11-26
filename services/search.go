package services

import (
	"strings"
	"unicode/utf8"
)

func CleanCity(q string) string {
	return strings.Join(strings.Fields(q), " ")
}

func ValidateCity(q string) string {
	length := utf8.RuneCountInString(q)
	if length == 0 {
		return "Veuillez saisir le nom d'une ville."
	}
	if length < 2 {
		return "Le nom doit contenir au moins 2 caractères."
	}
	if length > 60 {
		return "Le nom doit contenir au plus 60 caractères."
	}
	return ""
}
