package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Bayuaji64/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {

	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatal("Error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}

}
