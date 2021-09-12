package main

import (
	"context"
	"fmt"
	"log"

	"grpc_test/service"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("dial grpc server error: %v", err))
	}
	defer conn.Close()

	client := service.NewHelloServiceClient(conn)

	resp, err := client.Hello(context.TODO(), &service.Request{Name: "Bob"})

	log.Println(resp)

	log.Println("end...")
}
