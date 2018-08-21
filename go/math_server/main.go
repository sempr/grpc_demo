package main

import (
	"context"
	"log"
	"math"
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	pb "git.meideng.net/sempr/grpc-talks/go/math"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	if in.Value < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "说了不能为负数了")
	}
	return &pb.SqrtResponse{Value: math.Sqrt(in.Value)}, nil
}

func main() {
	addr := os.Getenv("GRPC_BIND")
	if addr == "" {
		addr = "0.0.0.0:50000"
	}

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})
	reflection.Register(s)
	log.Printf("Start listening on %s", addr)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
