package main

import (
	"log"

	pb "github.com/Bayuaji64/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

func main() {
	tls := true

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)

	// doLongGrret(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 1*time.Second)

}
