package helpers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type ErrorPage struct {
	Status int
	Message string
}

func Render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("template %q : %v", name, err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func RenderError(w http.ResponseWriter, path string, status int, msg string) {
    log.Printf("Erreur %v sur %v : %v", status, path, msg)   // dans ton terminal
    Render(w, status, "error", ErrorPage{Status: status, Message: msg})
}