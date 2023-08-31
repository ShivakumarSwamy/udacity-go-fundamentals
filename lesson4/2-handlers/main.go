package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var dictionary = map[string]string{
	"GO":     "A programming language created by Google.",
	"Gopher": "A Software engineer who builds with Go.",
	"GoLang": "Another name of Go.",
}

func main() {
	http.HandleFunc("/", getDictionary)

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func getDictionary(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(dictionary)
}
