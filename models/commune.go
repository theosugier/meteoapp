package models

type Commune struct {
	Code         string   `json:"code"`
	Nom          string   `json:"nom"`
	CodesPostaux []string `json:"codesPostaux"`
	Population   int      `json:"population"`
	Centre       Centre   `json:"centre"`
}

type Centre struct {
	Coordinates []float64 `json:"coordinates"`
}
