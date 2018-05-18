package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// test that the response handler returns a HTTP 200 status code
func TestResponseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/planeTracker", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(response)
	handler.ServeHTTP(recorder, req)
	status := recorder.Code

	if status != http.StatusOK {
		t.Errorf("Error. Got %v, want %v", status, http.StatusOK)
	}
}

// checks that getLat() returns correct value
func TestGetLat(t *testing.T) {
	coordinates := "(58.159912, 8.018206)"
	expected := "58.159912"
	res := getLat(coordinates)

	if res != expected {
		t.Errorf("Error. Got %v, want %v", res, expected)
	}
}

// checks that getLong() returns correct value
func TestGetLong(t *testing.T) {
	coordinates := "(58.159912, 8.018206)"
	expected := "8.018206"
	res := getLong(coordinates)

	if res != expected {
		t.Errorf("Error. Got %v, want %v", res, expected)
	}
}

// checks that joinUrl() returns correct string for API URL
func TestJoinUrl(t *testing.T) {
	lat := "58.159912"
	long := "8.018206"
	radius := "&fDstL=0&fDstU=100"
	expected := "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=58.159912&lng=8.018206&fDstL=0&fDstU=100"
	res := joinUrl(lat, long, radius)

	if res != expected {
		t.Errorf("Error. Got %v, want %v", res, expected)
	}
}

// checks that getApiUrl() returns correct string for API URL
func TestGetApiUrl(t *testing.T) {
	coordinates := "(58.159912, 8.018206)"
	expected := "http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=58.159912&lng=8.018206&fDstL=0&fDstU=100"
	res := getApiUrl(coordinates)

	if res != expected {
		t.Errorf("Error. Got %v, want %v", res, expected)
	}
}
