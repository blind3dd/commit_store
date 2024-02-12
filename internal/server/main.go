package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// NewHTTPServer creates a new HTTP Server
// it uses handlerFunc for index and v1 endpoint for
func NewHTTPServer(addr string) *http.Server {
	r := mux.NewRouter()
	httpSrv := newLogSrv()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode("{\"status\": \"OK\"}"); err != nil {
			log.Fatalf("failed to reach server, error: %v", err)
		}
	})
	// same endpoint for now to test mutex
	r.HandleFunc("/v1/commit", httpSrv.handleConsume).Methods("GET")
	r.HandleFunc("/v1/commit", httpSrv.handleProduce).Methods("POST")
	//TODO seperate producer from the consumer with gRPC or (not) whatever.
	// use api-gw annotation for proto3
	//r.Handle("/v1/produce", r.Methods("POST").GetHandler())
	//r.Handle("/v1/consume", r.Methods("GET").GetHandler())

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
