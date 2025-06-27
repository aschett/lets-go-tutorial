package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from andibox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form to create new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("snippet/view", snippetView)
	mux.HandleFunc("snippet/create", snippetCreate)

	log.Print("Starting Server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
