package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Aplicacao exemplo - 2.0")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
