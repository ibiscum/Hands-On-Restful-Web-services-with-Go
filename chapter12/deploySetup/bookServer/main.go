package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Book holds data of a book
type Book struct {
	ID            int
	ISBN          string
	Author        string
	PublishedYear string
}

func main() {
	// File open for reading, writing and appending
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	// This attaches program logs to file
	log.SetOutput(f)

	// Function handler for handling requests
	http.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%q", r.UserAgent())
		// Fill the book details
		book := Book{
			ID:            123,
			ISBN:          "0-201-03801-3",
			Author:        "Donald Knuth",
			PublishedYear: "1968",
		}
		// Convert struct to JSON using Marshal
		jsonData, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(jsonData)
		if err != nil {
			log.Fatal(err)
		}
	})
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
