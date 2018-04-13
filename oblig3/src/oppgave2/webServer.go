package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

func main() {

	http.HandleFunc("/", response)
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

type PageVariables struct {
	Name   string
	Number string
}

func response(w http.ResponseWriter, r *http.Request) {

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
		HomePageVars := PageVariables{
			Name:   v.Name,
			Number: v.Number,
		}
		err = t.Execute(w, HomePageVars) //fyrer av template og sender ved HomePageVars
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, v := range county.Entries {
		fmt.Println("Name: ", v.Name)
		fmt.Println("Name type = ", reflect.TypeOf(v.Name))

		fmt.Println("Number: ", v.Number)
	}

	/*county := Entries{}
	jsonErr := json.Unmarshal(body, &county)*/

	/*output, err := json.Marshal(county)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(output)*/

}
