package controllers

import (
	"meteoapp/helpers"
	"net/http"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "index", nil)
}

func Ville(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "ville", nil)
}

func Meteo(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, 200, "meteo", nil)
}

func PageInconnue(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, r.URL.Path, 404, "Cette page n'existe pas.") // r.URL.Path = le chemin tapé
}
