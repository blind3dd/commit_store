package main

import (
	"github.com/blind3dd/commit_store/internal/server"
	"log"
)

func main() {

	if err := server.NewHTTPServer(":8080").ListenAndServe(); err != nil {
		log.Fatalf("failed to listen, error: %v", err)
	}

	//srv := server.NewHTTPServer(":8088")
	//if err := srv.ListenAndServe(); err != nil {
	//	log.Fatalf("failed to listen, error: %v", err)
	//} else {
	//	log.Println("server has started, waiting to handle requests")
	//}
	//log.Fatal(srv.ListenAndServe())

}
