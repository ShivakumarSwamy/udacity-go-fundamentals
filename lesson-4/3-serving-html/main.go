package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Server starting at port 3000...")
	http.ListenAndServe(":3000", nil)
}
