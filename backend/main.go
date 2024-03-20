package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/api/v1/download", downloadFileHandler).Methods("GET")
	r.HandleFunc("/api/v1/upload", uploadFileHandler).Methods("POST")
	r.HandleFunc("/api/v1/process", processHandler).Methods("POST")

	// CORS middleware
	cors := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Proceed to the next middleware or handler
			next.ServeHTTP(w, r)
		})
	}

	srv := &http.Server{
		Handler: cors(r),
		Addr:    ":9090",
	}
	log.Fatal(srv.ListenAndServe())
}
