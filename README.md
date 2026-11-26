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
- /ville?q=Aix-en-Provence : une seule commune porte ce nom, on va direct sur sa meteo
- /ville?q=Saint : la liste des communes qui contiennent ce nom
- /ville?q=Saint-Martin : page de choix entre les communes qui ont le meme nom
- /ville?q= : erreur 400, il faut saisir une ville
- /ville?q=Zzzzz : erreur 404, aucune commune trouvee
- /meteo?code=13001 : la page d'Aix-en-Provence avec la carte (previsions pas encore finies)
- /meteo?code=99999 : erreur 404, la commune existe pas
- /page-inconnue : erreur 404

## Ce qui marche pour l'instant

- les routes et la page 404
- les templates avec le head et le footer communs
- le CSS dans assets/
- le formulaire de recherche et la verification de la saisie
- la recherche des communes avec l'API (et la page de choix si plusieurs ont le meme nom)
- les erreurs 400, 404 et 503 sur la recherche
- la page meteo d'une commune avec la carte OpenStreetMap

Reste a faire : les previsions des 7 jours (Open-Meteo) et les barres de temperature.

## Sources

- API Decoupage administratif (geo.api.gouv.fr)
- Open-Meteo.com (licence CC BY 4.0)
- OpenStreetMap
