package main

import (
	"log"
	"meteoapp/routes"
	"net/http"
)

func main() {
	log.Println("Serveur sur http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", routes.New())
	if err != nil {
		log.Fatal(err)
	}
}
