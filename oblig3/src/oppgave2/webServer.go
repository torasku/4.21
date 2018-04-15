package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

type Entry struct {
	Entries []struct {
		Name   string `json:"navn"`
		Number string `json:"nummer"`
	} `json:"entries"`
}

type PageVar1 struct {
	Name   string
	Number string
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

type PageVar2 struct {
	Latitude  string
	Name      string
	Plastic   string
	GlasMetal string
	Shoe      string
	Longitude string
}

var station Station

func response1(w http.ResponseWriter, r *http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/fylke"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, getErr := client.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(result.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	county := Entry{}
	jsonErr := json.Unmarshal(body, &county)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("index.html") //parse index.html
	if err != nil {
		log.Print(err)
	}

	for _, v := range county.Entries {
		HomePageVars := PageVar1{
			Name:   v.Name,
			Number: v.Number,
		}
		err = t.Execute(w, HomePageVars) //fyrer av template og sender ved HomePageVars
		if err != nil {
			log.Fatal(err)
		}
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
