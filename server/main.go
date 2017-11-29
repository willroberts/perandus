package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/willroberts/perandus/server/api"
)

var (
	bindAddress string = "0.0.0.0:8000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.StreamHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         bindAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting HTTP service")
	log.Fatal(srv.ListenAndServe())
}
