package main

import (
	"context"
	"log"
	"net"
	pb "statis/checkpoint"

	"google.golang.org/grpc"
)

const (
	port = ":1337"
)

type server struct {
	pb.UnimplementedCheckpointServer
}

func (s *server) CreateCheckpoint(ctx context.Context, in *pb.CheckpointRequest) (*pb.CheckpointReply, error) {
	log.Printf("Received Checkpoint: %v %v", in.GetName(), in.GetState())
	return &pb.CheckpointReply{Status: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCheckpointServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
