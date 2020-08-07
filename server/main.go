package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "net"
    "server_main_land/proto"
)

type server struct{}

func main() {
    listner, err := net.Listen("tcp", ":4040")
    if err != nil {
        panic(err)
    }

    srv := grpc.NewServer()
    proto.RegisterAddServiceServer(srv, &server{})
    reflection.Register(srv)

    fmt.Println("Listening on localhost:4040 using gRpc")
    if e := srv.Serve(listner); e != nil {
        panic(e)
    }
}


func (s *server) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
    fmt.Println("Calling add ======= >>")
    a, b := req.GetA(), req.GetB()
    result := a + b
    return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
    fmt.Println("Calling multiply ======= >>")
    a, b := req.GetA(), req.GetB()
    result := a * b
    return &proto.Response{Result: result}, nil
}