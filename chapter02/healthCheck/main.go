package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	_, err := io.WriteString(w, currentTime.String())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
