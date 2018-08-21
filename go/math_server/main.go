package main

import (
	"context"
	"io"
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

func (s *server) Stat(stream pb.Math_StatServer) error {
	var sum, count int32 = 0, 0
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.StatResponse{Sum: sum, Count: count})
			return nil
		}
		count++
		sum += in.Value
	}
}

func (s *server) Factor(in *pb.FactorRequest, stream pb.Math_FactorServer) error {
	val := in.Value
	var i int32
	for i = 2; i*i < val; i++ {
		for ; val%i == 0; val /= i {
			stream.Send(&pb.FactorResponse{Value: i})
		}
	}
	if val > 1 {
		stream.Send(&pb.FactorResponse{Value: val})
	}
	return nil
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
