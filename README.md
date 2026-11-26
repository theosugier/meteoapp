# meteoapp

TP Ynov : application meteo en Go.

On tape le nom d'une ville, le serveur trouve ses coordonnees avec l'API du gouvernement
(geo.api.gouv.fr) puis recupere la meteo des 7 prochains jours avec Open-Meteo.
Tout est fait en Go, pas de JavaScript.

## Lancer le projet

Il faut Go 1.22 minimum.

    git clone https://github.com/theosugier/meteoapp.git
    cd meteoapp
    go run .

Ensuite aller sur http://localhost:8080

## Exemples d'URL

- / : la page d'accueil avec le formulaire
- /ville?q=Marseille : la liste des communes trouvees
- /ville?q=Saint-Martin : page de choix entre les communes qui ont le meme nom
- /ville?q= : erreur 400, il faut saisir une ville
- /meteo?code=13001 : la meteo d'Aix-en-Provence (pas encore fini)
- /page-inconnue : erreur 404

## Ce qui marche pour l'instant

- les routes et la page 404
- les templates avec le head et le footer communs
- le CSS dans assets/
- le formulaire de recherche et la verification de la saisie
- la recherche des communes avec l'API

Reste a faire : les previsions meteo, la carte, les barres de temperature et les autres erreurs.

## Sources

- API Decoupage administratif (geo.api.gouv.fr)
- Open-Meteo.com (licence CC BY 4.0)
- OpenStreetMap
