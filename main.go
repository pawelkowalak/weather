package main

import (
	"fmt"
	"log"
	"net/http"
)

type weather struct {
	Temp  int  `json:"temp"`
	Windy bool `json:"windy"`
}

var city2weather = map[string]weather{}
var id int

func weatherHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello world from service #%d", id)
}

func healthHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "OK")
}

func main() {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// id = r.Intn(256)
	log.Printf("Service #%d Listening for connections", id)
	http.HandleFunc("/weather/", weatherHandler)
	http.HandleFunc("/health/", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
