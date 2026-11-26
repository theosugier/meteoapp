package controllers

import (
	"errors"
	"meteoapp/helpers"
	"meteoapp/models"
	"meteoapp/services"
	"net/http"
	"strings"
	"unicode/utf8"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "index", models.PageRecherche{})
}

func Ville(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q == "" {
		helpers.Render(w, 400, "index", models.PageRecherche{Valeur: q, Erreur: "Veuillez saisir le nom d'une ville."})
		return
	}
	if utf8.RuneCountInString(q) < 2 || utf8.RuneCountInString(q) > 60 {
		helpers.Render(w, 400, "index", models.PageRecherche{Valeur: q, Erreur: "Le nom doit contenir entre 2 et 60 caractères."})
		return
	}
	data, dataErr := services.SearchCommuneByName(q)
	if dataErr != nil {
		if errors.Is(dataErr, models.BadRequestError) {
			helpers.RenderError(w, "/ville", http.StatusBadRequest, "demande invalide")
			return
		}
		helpers.RenderError(w, r.URL.Path, 503, "Service momentanément indisponible.")
		return
	}
	if len(data) == 0 {
		helpers.RenderError(w, r.URL.Path, 404, "Aucune commune ne correspond à « "+q+" ».")
		return
	}
	helpers.Render(w, 200, "ville", data)
}

func Meteo(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "meteo", nil)
}

func PageInconnue(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, r.URL.Path, 404, "Cette page n'existe pas.") // r.URL.Path = le chemin tapé
}
