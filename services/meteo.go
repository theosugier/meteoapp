package services

import (
	"fmt"
	"meteoapp/api"
	"meteoapp/models"
)

// ecart autour de la commune pour la carte (donne par le sujet)
const ecartCarte = 0.05

// je prepare tout ce dont la page meteo a besoin
func PreparerPageMeteo(code string) (models.PageMeteo, error) {
	commune, err := api.CommuneParCode(code)
	if err != nil {
		return models.PageMeteo{}, err
	}

	// securite : sans les 2 coordonnees, impossible de placer la carte
	if len(commune.Centre.Coordinates) < 2 {
		return models.PageMeteo{}, models.InternalError
	}

	return models.PageMeteo{
		Commune:  commune,
		CarteURL: urlCarte(commune),
	}, nil
}

// l'URL de la carte OpenStreetMap, calculee en Go
func urlCarte(c models.Commune) string {
	lon := c.Centre.Coordinates[0] // attention : l'API Geo donne [longitude, latitude]
	lat := c.Centre.Coordinates[1]
	d := ecartCarte

	return fmt.Sprintf(
		"https://www.openstreetmap.org/export/embed.html?bbox=%.4f,%.4f,%.4f,%.4f&layer=mapnik&marker=%.4f,%.4f",
		lon-d, lat-d, lon+d, lat+d, lat, lon,
	)
}
