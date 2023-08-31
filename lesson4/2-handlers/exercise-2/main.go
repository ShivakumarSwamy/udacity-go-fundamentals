package main

import (
	"encoding/json"
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
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(cityPopulations)
	writer.WriteHeader(http.StatusOK)
}
