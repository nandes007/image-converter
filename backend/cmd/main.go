package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandes007/image-converter/internal"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", internal.IndexHandler).Methods("GET")
	r.HandleFunc("/api/v1/process", internal.ProcessHandler).Methods("POST")

	// CORS middleware
	cors := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}

	srv := &http.Server{
		Handler: cors(r),
		Addr:    ":9090",
	}
	log.Fatal(srv.ListenAndServe())
}
