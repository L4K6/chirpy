package main

import (
	"log"
	"net/http"
)

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal("problem writing the text body")
	}
}
