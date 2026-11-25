package routes

import (
	"meteoapp/controllers"
	"net/http"
)

func New() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", controllers.Accueil)
	router.HandleFunc("GET /ville", controllers.Ville)
	router.HandleFunc("GET /meteo", controllers.Meteo)

	fs := http.FileServer(http.Dir("assets"))                       // chercher dans dossier
	router.Handle("GET /static/", http.StripPrefix("/static/", fs)) // retire  /static

	router.HandleFunc("/", controllers.PageInconnue)

	return router
}
