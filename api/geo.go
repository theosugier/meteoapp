package api

import (
	"encoding/json"
	"fmt"
	"meteoapp/models"
	"net/http"
	"time"
	"net/url"
)

const geoURL = "https://geo.api.gouv.fr"

var client = &http.Client{Timeout: 5 * time.Second}

func ChercherCommunes(nom string) ([]models.Commune, error) {
	endpoint := fmt.Sprintf("%s/communes?nom=%s&fields=code,nom,codesPostaux,population,centre&boost=population&limit=10", geoURL, url.QueryEscape(nom))

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Erreur (%w) : Preparation de la requete, %v", models.InternalError, err)
	}

	request.Header.Set("User-Agent", "meteoapp-theosugier/1.0")

	response, responseErr := client.Do(request)
	if responseErr != nil {
		return nil, fmt.Errorf("Erreur (%w) : Envoie de la requete, %v", models.InternalError, responseErr)
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:

	case http.StatusNotFound:
		return nil, models.NotFoundError
	case http.StatusBadRequest:
		return nil, models.BadRequestError
	default:
		return nil, models.InternalError
	}

	var data []models.Commune
	errDecode := json.NewDecoder(response.Body).Decode(&data)
	if errDecode != nil {
		return nil, models.InternalError
	}
	return data, nil
}

func CommuneParCode(code string) (models.Commune, error) {
	endpoint := fmt.Sprintf("%s/communes/%s?fields=code,nom,codesPostaux,population,centre", geoURL, url.PathEscape(code))

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return models.Commune{}, fmt.Errorf("Erreur (%w) : Preparation de la requete, %v", models.InternalError, err)
	}

	request.Header.Set("User-Agent", "meteoapp-theosugier/1.0")

	response, responseErr := client.Do(request)
	if responseErr != nil {
		return models.Commune{}, fmt.Errorf("Erreur (%w) : Envoie de la requete, %v", models.InternalError, responseErr)
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:

	case http.StatusNotFound:
		return models.Commune{}, models.NotFoundError
	case http.StatusBadRequest:
		return models.Commune{}, models.NotFoundError
	default:
		return models.Commune{}, models.InternalError
	}

	var data models.Commune
	errDecode := json.NewDecoder(response.Body).Decode(&data)
	if errDecode != nil {
		return models.Commune{}, models.InternalError
	}
	return data, nil
}
