package controllers

import (
	"errors"
	"meteoapp/helpers"
	"meteoapp/models"
	"meteoapp/services"
	"net/http"
	"strings"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "index", models.PageRecherche{})
}

func Ville(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	data, dataErr := services.SearchCommuneByName(q)
	if dataErr != nil {
		if errors.Is(dataErr, models.BadRequestError) {
			helpers.RenderError(w, "/ville", http.StatusBadRequest, "demande invalide")
			return
		}
		//...
	}
	helpers.Render(w, 200, "ville", data)
}

func Meteo(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "meteo", nil)
}

func PageInconnue(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, r.URL.Path, 404, "Cette page n'existe pas.") // r.URL.Path = le chemin tapé
}
