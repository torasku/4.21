package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/planeTracker", response)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type Long struct {
	Long float64
}

type Lat struct {
	Lat float64
}

type Plane struct {
	Src   int `json:"src"`
	Feeds []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		PolarPlot bool   `json:"polarPlot"`
	} `json:"feeds"`
	SrcFeed int  `json:"srcFeed"`
	ShowSil bool `json:"showSil"`
	ShowFlg bool `json:"showFlg"`
	ShowPic bool `json:"showPic"`
	FlgH    int  `json:"flgH"`
	FlgW    int  `json:"flgW"`
	AcList  []struct {
		ID           int      `json:"Id"`
		Rcvr         int      `json:"Rcvr"`
		HasSig       bool     `json:"HasSig"`
		Icao         string   `json:"Icao"`
		Bad          bool     `json:"Bad"`
		Reg          string   `json:"Reg"`
		FSeen        string   `json:"FSeen"`
		TSecs        int      `json:"TSecs"`
		CMsgs        int      `json:"CMsgs"`
		Alt          int      `json:"Alt,omitempty"`
		GAlt         int      `json:"GAlt,omitempty"`
		InHg         float64  `json:"InHg"`
		AltT         int      `json:"AltT"`
		Call         string   `json:"Call,omitempty"`
		Lat          float64  `json:"Lat"`
		Long         float64  `json:"Long"`
		PosTime      int64    `json:"PosTime"`
		Mlat         bool     `json:"Mlat"`
		Tisb         bool     `json:"Tisb"`
		Spd          float64  `json:"Spd,omitempty"`
		Trak         float64  `json:"Trak"`
		TrkH         bool     `json:"TrkH"`
		Type         string   `json:"Type"`
		Mdl          string   `json:"Mdl"`
		Man          string   `json:"Man"`
		CNum         string   `json:"CNum"`
		Op           string   `json:"Op"`
		Sqk          string   `json:"Sqk"`
		Help         bool     `json:"Help,omitempty"`
		VsiT         int      `json:"VsiT"`
		Dst          float64  `json:"Dst"`
		Brng         float64  `json:"Brng"`
		WTC          int      `json:"WTC"`
		Species      int      `json:"Species"`
		Engines      string   `json:"Engines"`
		EngType      int      `json:"EngType"`
		EngMount     int      `json:"EngMount"`
		Mil          bool     `json:"Mil"`
		Cou          string   `json:"Cou"`
		HasPic       bool     `json:"HasPic"`
		Interested   bool     `json:"Interested"`
		FlightsCount int      `json:"FlightsCount"`
		Gnd          bool     `json:"Gnd,omitempty"`
		SpdTyp       int      `json:"SpdTyp"`
		CallSus      bool     `json:"CallSus"`
		Trt          int      `json:"Trt"`
		OpIcao       string   `json:"OpIcao,omitempty"`
		Year         string   `json:"Year,omitempty"`
		Sig          int      `json:"Sig,omitempty"`
		PosStale     bool     `json:"PosStale,omitempty"`
		Vsi          int      `json:"Vsi,omitempty"`
		PicX         int      `json:"PicX,omitempty"`
		PicY         int      `json:"PicY,omitempty"`
		From         string   `json:"From,omitempty"`
		To           string   `json:"To,omitempty"`
		Stops        []string `json:"Stops,omitempty"`
		TAlt         int      `json:"TAlt,omitempty"`
		TTrk         float64  `json:"TTrk,omitempty"`
		Tag          string   `json:"Tag,omitempty"`
	} `json:"acList"`
	TotalAc   int    `json:"totalAc"`
	LastDv    string `json:"lastDv"`
	ShtTrlSec int    `json:"shtTrlSec"`
	Stm       int64  `json:"stm"`
}

var planes Plane
var apiBase = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?"
var radius = "&fDstL=0&fDstU=100"

func response(w http.ResponseWriter, r *http.Request) {

	var url string
	r.ParseForm()
	coordinates := r.FormValue("coordinates")
	if len(coordinates) == 0 {
		fmt.Println("her")
		url = "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=58.159912&lng=8.018206&fDstL=0&fDstU=100"
	} else {
		url = getApiUrl(coordinates)
	}
	getData(url)

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, planes)
	if err != nil {
		log.Fatal(err)
	}
}

func getData(url string) {
	result, _ := http.Get(url)
	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &planes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

func getLat(coordinates string) string {
	trim := strings.Trim(coordinates, "()")
	split := strings.Split(trim, ",")
	lat := split[0]

	return lat
}

func getLong(coordinates string) string {
	trim := strings.Trim(coordinates, "()")
	split := strings.Split(trim, ",")
	long := split[1]
	long = strings.TrimSpace(long)

	return long
}

func joinUrl(lat string, lng string, r string) string {
	slice := []string{apiBase, lat}
	url := strings.Join(slice, "lat=")

	slice2 := []string{url, lng}
	url = strings.Join(slice2, "&lng=")

	slice3 := []string{url, r}
	url = strings.Join(slice3, "")

	return url
}

func getApiUrl(coordinates string) string {
	lat := getLat(coordinates)
	long := getLong(coordinates)

	apiUrl := joinUrl(lat, long, radius)

	return apiUrl
}
