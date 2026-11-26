package services

import (
	"meteoapp/api"
	"meteoapp/models"
)

func SearchCommuneByName(query string) ([]models.Commune,error){
	return api.ChercherCommunes(query)
}