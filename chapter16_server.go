// Chapter 16: gRPC Server
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "path/to/order"
)

type server struct {
    pb.UnimplementedOrderServiceServer
}

func (s *server) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    return &pb.OrderResponse{Success: true, Message: "Order placed"}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    grpcServer := grpc.NewServer()
    pb.RegisterOrderServiceServer(grpcServer, &server{})
    log.Fatal(grpcServer.Serve(lis))
}
