package models

type PageRecherche struct {
	Valeur string
	Erreur string
}

// ce que la page meteo recoit
type PageMeteo struct {
	Commune  Commune
	CarteURL string
}
