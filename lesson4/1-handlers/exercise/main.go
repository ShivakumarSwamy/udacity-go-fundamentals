package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/citylist", cityList)

	fmt.Printf("Server is starting...")
	http.ListenAndServe(":3000", nil)

}

func cityList(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "List of most populous cities:\n")

	for _, city := range cities {
		fmt.Fprintf(writer, "%s\n", city)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Main Page")

}
