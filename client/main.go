package main

import (
	"context"
	"log"

	"github.com/renanbastos93/poc-grpc/pb"
	"google.golang.org/grpc"
)

// Hello ...
func Hello(err error, client pb.HelloServiceClient) {
	request := &pb.HelloRequest{
		Name: "Renan Bastos",
	}

	response, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalln("[Hello()]", err.Error())
	}
	log.Println("Response: ", response.Msg)
}

func main() {
	var (
		client pb.HelloServiceClient
		conn   *grpc.ClientConn
		err    error
	)

	if conn, err = grpc.Dial("0.0.0.0:50051", grpc.WithInsecure()); err != nil {
		log.Fatalln("[DIAL]", err.Error())
	}

	defer func() {
		_ = conn.Close()
	}()

	client = pb.NewHelloServiceClient(conn)
	Hello(err, client)
}
