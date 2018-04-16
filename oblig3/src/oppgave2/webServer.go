package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", response1)
	r.HandleFunc("/2", response2)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// lager structs for aa holde paa JSON data

type County struct {
	Entries []struct {
		Name   string `json:"navn"`
		Number string `json:"nummer"`
	} `json:"entries"`
}

type Station struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Name      string `json:"navn"`
		Plastic   string `json:"plast"`
		GlasMetal string `json:"glass_metall"`
		Shoe      string `json:"tekstil_sko,omitempty"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

var station Station
var county County

func response1(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/fylke"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &county)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("counties.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, county)
	if err != nil {
		log.Fatal(err)
	}
}

func response2(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &station)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("stations.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, station)
	if err != nil {
		log.Fatal(err)
	}

}
