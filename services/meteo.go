package services

import (
	"fmt"
	"meteoapp/api"
	"meteoapp/models"
)

const ecartCarte = 0.05

func PreparerPageMeteo(code string) (models.PageMeteo, error) {
	commune, err := api.CommuneParCode(code)
	if err != nil {
		return models.PageMeteo{}, err
	}

	if len(commune.Centre.Coordinates) < 2 {
		return models.PageMeteo{}, models.InternalError
	}

	return models.PageMeteo{
		Commune:  commune,
		CarteURL: urlCarte(commune),
	}, nil
}

func urlCarte(c models.Commune) string {
	lon := c.Centre.Coordinates[0]
	lat := c.Centre.Coordinates[1]
	d := ecartCarte

	return fmt.Sprintf(
		"https://www.openstreetmap.org/export/embed.html?bbox=%.4f,%.4f,%.4f,%.4f&layer=mapnik&marker=%.4f,%.4f",
		lon-d, lat-d, lon+d, lat+d, lat, lon,
	)
}
