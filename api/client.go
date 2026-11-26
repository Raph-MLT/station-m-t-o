package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

const userAgent = "station-meteo/1.0 (projet etudiant Ynov)"

var ErrNotFound = errors.New("ressource introuvable")

func getJSON(address string, target any) error {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s : code %d", address, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}
