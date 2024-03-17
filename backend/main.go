package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/download", downloadFileHandler).Methods("GET")
	r.HandleFunc("/upload", uploadFileHandler).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}
