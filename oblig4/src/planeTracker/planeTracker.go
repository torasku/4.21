package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/planeTracker", response)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
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

func response(w http.ResponseWriter, r *http.Request) {

	url := "https://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=58.14671&lng=7.9956&fDstL=0&fDstU=100"
	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &planes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, planes)
	if err != nil {
		log.Fatal(err)
	}
}
