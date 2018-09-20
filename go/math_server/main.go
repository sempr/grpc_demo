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

// UnaryServerInterceptor ....
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. Info: %+v", info)
	resp, err := handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

// StreamServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Streaming RPCs.
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("after handling. err: %v", err)
	return err
}

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

type server2 struct{}

func (s *server2) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{C: in.A + in.B}, nil
}

func (s *server2) Sub(ctx context.Context, in *pb.SubRequest) (*pb.SubResponse, error) {
	return &pb.SubResponse{C: in.A - in.B}, nil
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

	s := grpc.NewServer(grpc.StreamInterceptor(StreamServerInterceptor),
		grpc.UnaryInterceptor(UnaryServerInterceptor))

	pb.RegisterMathServer(s, &server{})
	pb.RegisterMath2Server(s, &server2{})
	reflection.Register(s)
	log.Printf("Start listening on %s", addr)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
