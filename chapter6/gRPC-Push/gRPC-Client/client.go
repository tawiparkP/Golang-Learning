package main

import (
	"io"
    "log"
    
    pb "github.com/tawiparkP/protofiles"
    "google.golang.org/grpc"
    "golang.org/x/net/context"

)

const (
    address = "localhost:50051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest){
    log.Println("Started Listening to the server")
    stream, err := client.MakeTransaction(context.Background(),request)
    if err != nil {
        log.Fatalf("%v.MakeTransaction(_) = _, %v", client,err)
    }

    for {
        response, err := stream.Recv()
        if err == io.EOF {
            break
        } 

        if err != nil {
            log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
        }

        log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
    }
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {

    }

    c := pb.NewMoneyTransactionClient(conn)

    from := "1234"
    to := "5678"
    amount := float32(1250.25)
    ReceiveStream(c,&pb.TransactionRequest{From: from, To: to, Amount: amount})
}



