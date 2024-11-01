package main

import (
	"log"

	pb "github.com/ibiscum/Hands-On-Restful-Web-services-with-Go/chapter06/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Frontend or App
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Contact the server and print out its response.
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from,
		To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}
