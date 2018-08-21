package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "git.meideng.net/sempr/grpc-talks/go/math"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	addr := os.Getenv("GRPC_SERVER")
	if addr == "" {
		addr = "127.0.0.1:50002"
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMathClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Sqrt(ctx, &pb.SqrtRequest{Value: 10.0})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%f", r.Value)
	}

	r, err = c.Sqrt(ctx, &pb.SqrtRequest{Value: -10.0})
	if err != nil {
		statusCode, _ := status.FromError(err)
		log.Println(statusCode.Code(), statusCode.Message())
	}

}
