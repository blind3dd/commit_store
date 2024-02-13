package codec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blind3dd/commit_store/internal/commit"
	"github.com/blind3dd/commit_store/internal/entities"
	"log"
	"net/http"
)

// HttpServer defines log records slice
type HttpServer struct {
	Commit *commit.Commit
}

func NewLogSrv() *HttpServer {
	return &HttpServer{
		Commit: commit.NewCommit(),
	}
}

// HandleProduce defines handler for Producer Request and Response
func (hs *HttpServer) HandleProduce(w http.ResponseWriter, r *http.Request) {
	var req entities.ProducerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity) // 422 for now
		log.Printf("failed to encode producer request, error: %v", err)
	}
	offset, err := hs.Commit.Write(req.Commit)
	if err != nil {
		log.Printf("failed to write commit log, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp := entities.ProducerResponse{Offset: offset}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode produdcer response, error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleConsume defines handler for Producer Request and Response
func (hs *HttpServer) HandleConsume(w http.ResponseWriter, r *http.Request) {
	var req entities.ConsumerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("failed to decode consumer request, error: %v", err)
		http.Error(w, err.Error(), http.StatusTeapot)
	}
	logRecord, err := hs.Commit.Read(req.Offset)
	if errors.Is(err, commit.OffsetNotFoundError) {
		log.Printf("failed to find log commit entry, error: %v", err)
		_ = fmt.Errorf("invalid offset defined")
	} else if err != nil {
		log.Printf("failed to read from server, error: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	resp := entities.ConsumerResponse{Commit: logRecord}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode consumer response, error: %v", err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

//
//TODO FetchAll, Header Check for io.EOF etc
//TODO gRPC implementation of the above
