package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	fmt.Printf("Starting server on port :8000\n")
	if error := http.ListenAndServe(":8000", nil); error != nil {
		log.Fatal(error)
	}
}
