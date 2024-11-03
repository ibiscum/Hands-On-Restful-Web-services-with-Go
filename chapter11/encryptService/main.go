package main

import (
	fmt "fmt"
	"log"

	proto "github.com/ibiscum/Hands-On-Restful-Web-services-with-Go/chapter11/encryptService/proto"
	micro "go-micro.dev/v5"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("encrypter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	err := proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))
	if err != nil {
		log.Fatal(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
