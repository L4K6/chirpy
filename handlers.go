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
	header.Set("Content-Type", "text/html; charset=utf-8")
	numOfHits := cfg.fileserverHits.Load()
	stringToRender := fmt.Sprintf(
		`
	<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
`, numOfHits)
	_, err := w.Write([]byte(stringToRender))
	if err != nil {
		log.Fatal("problem generating metrics html")
	}
}

func (cfg *apiConfig) resetCountHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}
