package server

import (
	"github.com/blind3dd/commit_store/internal/commit"
	"net/http"
)

type httpServer struct {
	Commit *commit.Commit
}

func newLogSrv() *httpServer {
	return &httpServer{
		Commit: commit.NewCommit(),
	}
}

// ProducerRequest defines REST API request that based on the offset
// returns proper record as per Read func in log.go
type ProducerRequest struct {
	Record commit.Record `json:"record"`
}

// ProducerResponse returns timestamp based on the value
// of the commited log entry
type ProducerResponse struct {
	Offset uint64 `json:"offset"`
}

// ConsumerRequest defines an offset
type ConsumerRequest struct {
	Offset uint64 `json:"offset"`
}

// ConsumerResponse sends record log
type ConsumerResponse struct {
	Record commit.Record `json:"record"`
}

func (hs *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
}

func (hs *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
}

//TODO GRPC implementation of the above
