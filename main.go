package main

import (
	"fmt"
	"log"
	"net/http"
)

const appStatus string = "ONLINE"

func main() {

	fmt.Println("Server started and listening")
	http.HandleFunc("/", basicRequest)
	log.Fatal(http.ListenAndServe(":9003", nil))
}

func basicRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, appStatus)
}
