package main

import (
	"context"
	"log"
	"time"

	proto "github.com/ibiscum/Hands-On-Restful-Web-services-with-Go/chapter11/asyncService/proto"
	micro "go-micro.dev/v4"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("weather"),
	)
	p := micro.NewEvent("alerts", service.Client())

	go func() {
		for now := range time.Tick(15 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			err := p.Publish(context.TODO(), &proto.Event{
				City:        "Munich",
				Timestamp:   now.UTC().Unix(),
				Temperature: 2,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	// Init will parse the command line flags.
	service.Init()

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
