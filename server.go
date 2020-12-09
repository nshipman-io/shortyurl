package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {

	router:= mux.NewRouter()

	router.HandleFunc("/", Home)
	router.HandleFunc("/api/shorturl/new", New)
	router.HandleFunc("/api/shorturl/{url_id}", Process)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: router,
	}
	log.Println("Starting server at addr:", server.Addr)
	log.Fatal(server.ListenAndServe())
}