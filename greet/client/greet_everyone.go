package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Bayuaji64/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {

	log.Println("doGreetEveryonet was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatal("Error while calling GreetEveryone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Bayuaji"},
		{FirstName: "Pradipta"},
		{FirstName: "Arinanda"},
	}

	waitc := make(chan struct{})

	go func() {

		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}

		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc

}
