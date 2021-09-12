package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc_test/protos"
	"grpc_test/service"

	"google.golang.org/grpc"
)

type helloService struct {
	service.UnimplementedHelloServiceServer
}

func (h *helloService) Hello(ctx context.Context, req *service.Request) (*protos.Hello, error) {
	if req == nil || "" == req.Name {
		return nil, fmt.Errorf("request is not ok: %v", req)
	}

	ret := protos.CreateHello(req.Name)
	log.Println("resp:", ret)

	return ret, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	service.RegisterHelloServiceServer(s, &helloService{})

	log.Println("serving...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("start....")
}
