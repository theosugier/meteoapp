package helpers

import (
	"html/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type ErrorPage struct {
	Status int
	Message string
}
