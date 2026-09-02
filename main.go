package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	reqMultiplexer := http.NewServeMux()
	server := http.Server{Addr: ":" + port, Handler: reqMultiplexer}
	log.Printf("Serving on port: %s", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

}
