package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the server!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
