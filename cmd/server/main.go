package main

import (
	"github.com/blind3dd/store_commit_poc/internal/server"
	"log"
)

func main() {

	if err := server.NewHTTPServer(":8080").ListenAndServe(); err != nil {
		log.Fatalf("failed to listen, error: %v", err)
	}
	//srv := server.NewHTTPServer(":8080")
	//log.Fatal(srv.ListenAndServe())

}
