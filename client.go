package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jordancalder/payments"

	"google.golang.org/grpc"
)

func printTransactions(client pb.TransactionClient) {
	empty := &pb.Empty{}
	log.Printf("printing transactions...")
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	stream, err := client.GetTransactionsStream(ctx, empty)
	if err != nil {
		log.Fatal("ohnosuddendeath")
	}
	for {
		transaction, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when calling transactions stream: %s", err)
		}
		log.Printf("transaction: %d %s", uint64(transaction.Price), transaction.Name)
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
