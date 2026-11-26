package models

type PageRecherche struct {
	Valeur string
	Erreur string
}

type PageMeteo struct {
	Commune  Commune
	CarteURL string
}
