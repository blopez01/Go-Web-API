package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Locations []Location

func handleRequests() {
	pageRouter := mux.NewRouter().StrictSlash(true)

	pageRouter.HandleFunc("/", homePage)
	pageRouter.HandleFunc("/all", returnAllLocations)
	pageRouter.HandleFunc("/location/", addLocation).Methods("POST")
	pageRouter.HandleFunc("/location/{name}", deleteLocation).Methods("DELETE")
	pageRouter.HandleFunc("/location/{name}", updateLocation).Methods("PUT")
	pageRouter.HandleFunc("/location/{name}", returnSingleLocation)

	log.Fatal(http.ListenAndServe(":8080", pageRouter))
}

func main() {
	fmt.Println("RuneScape Locations API")
	Locations = []Location{
		{Name: "Lumbridge", Kingdom: "Misthalin", NPCs: 46},
		{Name: "Varrock", Kingdom: "Misthalin", NPCs: 78},
		{Name: "Ardougne", Kingdom: "Kandarin", NPCs: 84},
	}
	handleRequests()
}
