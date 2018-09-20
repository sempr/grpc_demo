package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "git.meideng.net/sempr/grpc-talks/go/math"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryClientInterceptor xxx
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("before invoker. method: %+v, request:%+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("after invoker. reply: %+v", reply)
	return err
}

// StreamClientInterceptor yyy
func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("before invoker. method: %+v, StreamDesc:%+v", method, desc)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("before invoker. method: %+v", method)
	return clientStream, err
}

func main() {

	addr := os.Getenv("GRPC_SERVER")
	if addr == "" {
		addr = "127.0.0.1:50000"
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(UnaryClientInterceptor),
		grpc.WithStreamInterceptor(StreamClientInterceptor),
	)
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

	// stream code here

	stream, err := c.Stat(ctx)
	if err != nil {
		log.Fatal(err)
	}
	stream.Send(&pb.StatRequest{Value: 1})
	stream.Send(&pb.StatRequest{Value: 2})
	stream.Send(&pb.StatRequest{Value: 1})
	stream.Send(&pb.StatRequest{Value: 2})
	stream.Send(&pb.StatRequest{Value: 1})

	reply, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Count, reply.Sum)

	stream1, err := c.Factor(ctx, &pb.FactorRequest{Value: 1080})
	if err != nil {
		log.Fatal(err)
	}
	for {
		val, err := stream1.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(val.Value)
	}

}
