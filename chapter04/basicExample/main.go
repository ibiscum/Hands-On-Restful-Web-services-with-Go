package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func main() {
	// Create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write to the response
	_, err := io.WriteString(resp, time.Now().String())
	if err != nil {
		log.Fatal(err)
	}
}
