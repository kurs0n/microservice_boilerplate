package main

import (
	"context"
	pb2 "essa/gen/broker"
	pb "essa/gen/greeting"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

type brokerServer struct {
	pb2.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request for name: %s", req.Name)
	return &pb.HelloResponse{
		Message: "Hello, " + req.Name + "!",
	}, nil
}

func (s *brokerServer) SayHello2(ctx context.Context, req *pb2.HelloRequest) (*pb2.HelloResponse, error) {
	log.Printf("Received request for name: %s", req.Name)
	return &pb2.HelloResponse{
		Message: "Hello, " + req.Name + "!",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})
	pb2.RegisterGreeterServer(grpcServer, &brokerServer{})
	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
