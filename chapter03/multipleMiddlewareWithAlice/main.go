package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

// Middleware to check content type as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Your resource creation logic goes here. For now it is plain print to console
		log.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		// Tell everything is fine
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("201 - Created"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 - Method Not Allowed"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
