package main

import (
	"context"
	"log"
	"net"

	"github.com/renanbastos93/poc-grpc/pb"
	"google.golang.org/grpc"
)

type server struct{}

// Hello ...
// Hello(context.Context, *HelloRequest) (*HelloResponse, error)
func (srv *server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	res := &pb.HelloResponse{
		Msg: "Hello " + request.GetName(),
	}
	return res, nil
}

func main() {
	var (
		ln         net.Listener
		err        error
		gRPCServer *grpc.Server
	)
	if ln, err = net.Listen("tcp", "0.0.0.0:50051"); err != nil {
		log.Fatalln("[LN]: ", err.Error())
	}
	gRPCServer = grpc.NewServer()
	pb.RegisterHelloServiceServer(gRPCServer, &server{})

	if err = gRPCServer.Serve(ln); err != nil {
		log.Fatalln("[gRpc Serve()]", err)
	}
}
