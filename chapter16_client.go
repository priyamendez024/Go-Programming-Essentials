// Chapter 16: gRPC Client
package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "path/to/order"
)

func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    defer conn.Close()
    client := pb.NewOrderServiceClient(conn)
    res, err := client.PlaceOrder(context.Background(), &pb.OrderRequest{Id: 123, Items: []string{"apple", "banana"}})
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Response:", res.Message)
}
