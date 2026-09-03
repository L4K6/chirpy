package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := http.Server{Addr: ":" + port, Handler: mux}
	log.Printf("Serving on port: %s", port)

	fileServerHandler := http.FileServer(http.Dir("."))
	fileServerHandler = http.StripPrefix("/app", fileServerHandler)

	var apiCfg apiConfig

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fileServerHandler))
	mux.HandleFunc("GET /api/healthz", endpointHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.reqCountHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetCountHandler)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
