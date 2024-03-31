package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Bayuaji64/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("GreetWithDeadline function was invoked")

	for i := 0; i < 3; i++ {

		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("The Client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)

	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
