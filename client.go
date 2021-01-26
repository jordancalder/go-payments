package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "./payments/payments.pb.go"

	"google.golang.org/grpc"
)

func printTransactions(client pb.TransactionClient) {
	log.Printf("printing transactions...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetTransactionsStream()
	if err != nil {
		log.Fatal("didnotdie")
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("holysmokes")
		}
		log.Printf(feature)
	}
}

func main() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("burrrrn")
	}
	defer conn.Close()
	client := pb.NewTransactionClient(conn)

	printTransactions(client)
}
