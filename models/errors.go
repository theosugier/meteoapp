package models

import "errors"

var (
	NotFoundError   = errors.New("ressource non trouvé")
	BadRequestError = errors.New("requete invalide")
	InternalError   = errors.New("erreur interne")
)
