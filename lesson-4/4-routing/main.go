package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var dictionary = map[string]string{
	"GO":     "A programming language created by Google.",
	"Gopher": "A Software engineer who builds with Go.",
	"GoLang": "Another name of Go.",
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getDictionary).Methods("GET")
	router.HandleFunc("/", create).Methods("POST")
	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}

func create(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newEntry map[string]string

	requestBody, _ := io.ReadAll(request.Body)

	json.Unmarshal(requestBody, &newEntry)

	for k, v := range newEntry {
		if _, ok := dictionary[k]; ok {
			writer.WriteHeader(http.StatusConflict)
		} else {
			dictionary[k] = v
			writer.WriteHeader(http.StatusCreated)
		}
	}
	json.NewEncoder(writer).Encode(dictionary)
}

func getDictionary(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dictionary)
}
