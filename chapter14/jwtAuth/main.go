package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/gorilla/mux"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))
var users = map[string]string{"naren": "passme", "admin": "password"}

// Response is a representation of JSON response for JWT
type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

// HealthcheckHandler returns the date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
	token, err := jwt.Parse(string(buf), func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("Access Denied; Please check the access token"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// If token is valid
		response := make(map[string]string)
		// response["user"] = claims["username"]
		response["time"] = time.Now().String()
		response["user"] = claims["username"].(string)
		responseJSON, _ := json.Marshal(response)
		_, err := w.Write(responseJSON)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// LoginHandler validates the user credentials
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			// Create a claims map
			claims := jwt.MapClaims{
				"username":  username,
				"ExpiresAt": 15000,
				"IssuedAt":  time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					log.Fatal(err)
				}
			}
			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(responseJSON)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getToken", getTokenHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
