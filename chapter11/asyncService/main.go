package main

import (
	"context"
	"log"
	"time"

	micro "github.com/go-micro/go-micro"
	proto "github.com/ibiscum/Hands-On-Restful-Web-services-with-Go/chapter11/asyncService/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("weather"),
	)
	p := micro.NewPublisher("alerts", service.Client())

	go func() {
		for now := range time.Tick(15 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			p.Publish(context.TODO(), &proto.Event{
				City:        "Munich",
				Timestamp:   now.UTC().Unix(),
				Temperature: 2,
			})
		}
	}()
	// Init will parse the command line flags.
	service.Init()

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
