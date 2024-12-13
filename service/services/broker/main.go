package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "essa/gen/broker"

	"google.golang.org/grpc"
)

var client pb.GreeterClient

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello2(ctx, &pb.HelloRequest{Name: "test"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	w.Write([]byte(response.Message))
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client = pb.NewGreeterClient(conn)

	http.HandleFunc("/hello", helloHandler)

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
