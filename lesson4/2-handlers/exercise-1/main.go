package main

import (
	"fmt"
	"net/http"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func main() {
	http.HandleFunc("/", getCityPopulations)
	fmt.Println("Server starting at port 3000...")
	http.ListenAndServe(":3000", nil)
}

func getCityPopulations(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	for city, population := range cityPopulations {
		fmt.Fprintf(writer, "<h2> City: %s; Population: %d</h2>\n", city, population)
	}
	writer.WriteHeader(http.StatusOK)
}
