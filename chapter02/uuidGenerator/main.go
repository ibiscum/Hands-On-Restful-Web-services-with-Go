package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

// UUID is a custom multiplexer
type UUID struct {
}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w)
		return
	}
	http.NotFound(w, r)
}

func giveRandomUUID(w http.ResponseWriter) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	tmp := fmt.Sprintf("%x", b)
	fmt.Fprint(w, tmp)
}

func main() {
	mux := &UUID{}
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
