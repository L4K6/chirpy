package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := http.Server{Addr: ":" + port, Handler: mux}
	log.Printf("Serving on port: %s", port)

	fileServerHandler := http.FileServer(http.Dir("."))
	fileServerHandler = http.StripPrefix("/app", fileServerHandler)

	mux.Handle("/app/", fileServerHandler)
	mux.HandleFunc("/healthz", EndpointHandler)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

}
