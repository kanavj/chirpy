package main

import (
	"log"
	"net/http"

	"github.com/crayboi420/chirpy/internal/database"
)

func main() {
	const filepathRoot = "./files/"
	const port = "8080"

	db, _ := database.NewDB("./database.json")
	apiCfg := apiConfig{
		fileserverHits: 0,
		db:             *db,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", healthz)
	mux.HandleFunc("GET /admin/metrics", apiCfg.middlewareMetricsHits)
	mux.HandleFunc("/api/reset", apiCfg.middlewareMetricsReset)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirpsCreate)
	mux.HandleFunc("GET /api/chirps/", apiCfg.handlerChirpsRetrieve)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerChirpRetrieve)
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("GET /api/users", apiCfg.handlerUsersRetrieve)
	mux.HandleFunc("GET /api/users/{userID}", apiCfg.handlerUserRetrieve)

	corsMux := middlewareCORS(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
