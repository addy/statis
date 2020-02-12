package main

import (
	"context"
	"log"
	"os"
	pb "statis/checkpoint"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:1337"
	defaultName = "DEFAULT_NAME"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCheckpointClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateCheckpoint(ctx, &pb.CheckpointRequest{Name: name, Timestamp: time.Now().UnixNano() / int64(time.Millisecond)})

	if err != nil {
		log.Fatalf("Could not create checkpoint: %v", err)
	}

	log.Printf("Status: %d", r.GetStatus())
}
