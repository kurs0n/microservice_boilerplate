package main

import (
	"context"
	"log"
	"time"

	pb "essa/gen/greeting"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Hackathon"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
