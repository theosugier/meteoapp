package api

import (
	"encoding/json"
	"fmt"
	"meteoapp/models"
	"net/http"
	"time"
)

const geoURL = "https://geo.api.gouv.fr"

var client = &http.Client{Timeout: 5 * time.Second}

func ChercherCommunes(nom string) ([]models.Commune, error) {
	endpoint := fmt.Sprintf("%s/communes?nom=%s&fields=code,nom,codesPostaux,population,centre", geoURL, nom)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Erreur (%w) : Preparation de la requete, %v", models.InternalError, err)
	}

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
