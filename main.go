package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Location struct {
	Name    string `json:"Name"`
	Kingdom string `json:"Kingdom"`
	NPCs    int    `json:"NPCs"`
}

var Locations []Location

func returnAllLocations(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPlayers")
	json.NewEncoder(writer).Encode(Locations)
}
func returnSingleLocation(writer http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)
	name := requestVars["name"]

	for _, location := range Locations {
		if strings.EqualFold(location.Name, name) {
			json.NewEncoder(writer).Encode(location)
		}
	}
}
func addLocation(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var location Location
	json.Unmarshal(reqBody, &location)
	Locations = append(Locations, location)

	json.NewEncoder(writer).Encode(location)
}
func deleteLocation(writer http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)
	name := requestVars["name"]

	for index, location := range Locations {
		if strings.EqualFold(location.Name, name) {
			Locations = append(Locations[:index], Locations[index+1:]...)
		}
	}
}
func updateLocation(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	var updatedLocation Location
	reqBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(reqBody, &updatedLocation)
	for i, location := range Locations {
		if strings.EqualFold(location.Name, name) {
			location.Name = updatedLocation.Name
			location.Kingdom = updatedLocation.Kingdom
			location.NPCs = updatedLocation.NPCs
			Locations[i] = location
			json.NewEncoder(writer).Encode(location)
		}
	}
}
func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Homepage")
	fmt.Println("Endpoint Hit: homePage")
}
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
