package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	err := http.ListenAndServe(":8000", loggedRouter)
	if err != nil {
		log.Fatal(err)
	}
}
