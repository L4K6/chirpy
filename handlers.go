package main

import (
	"fmt"
	"log"
	"net/http"
)

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal("problem writing the text body")
	}
}

func (cfg *apiConfig) reqCountHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	numOfHits := cfg.fileserverHits.Load()
	fmt.Fprintf(w, "Hits: %d", numOfHits)
}

func (cfg *apiConfig) resetCountHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}
