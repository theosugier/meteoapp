package services

import (
	"meteoapp/api"
	"meteoapp/models"
	"strings"
)

func SearchCommuneByName(query string) ([]models.Commune,error){
	return api.ChercherCommunes(query)
}

func CommuneUnique(communes []models.Commune, nom string) (models.Commune, bool) {
	var trouvee models.Commune
	nb := 0
	for _, c := range communes {
		if strings.EqualFold(c.Nom, nom) {
			trouvee = c
			nb++
		}
	}
	return trouvee, nb == 1
}
