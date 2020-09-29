package main

import (
	"net"
	"fmt"
	"time"
    "log"
    
    pb "github.com/tawiparkP/protofiles"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    port = ":50051"
    noOfStep = 3
)

type server struct{}

func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
    log.Printf("Got request for money transfer...")
    log.Printf("Amount: $%f, From A/c: %s, To A/c: %s", in.Amount, in.From, in.To)
    for i:= 0; i < noOfStep; i++{
        time.Sleep(time.Second*10)
        if err := stream.Send(&pb.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintf("Performing step %d", int32(i))}); err != nil {
            log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
        }
    }
    log.Printf("Successfully transaction amount $%v from %v to %v", in.Amount, in.From, in.To)
    return nil
}

func main(){
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterMoneyTransactionServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}