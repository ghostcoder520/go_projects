package main

import (
	"context"
	"fmt"
	pb "grpc-demo/hello-server/proto"
	"net"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.UnimplementedSayHelloServer
}

func (s helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("hello~" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello~" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")

	// 创建grpc服务
	grpcServer := grpc.NewServer()

	// 在grpc服务中注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &helloServer{})

	// 启动grpc服务
	err := grpcServer.Serve(listen)

	if err != nil {
		fmt.Println("failed to serve: %v", err)

		return
	}

}
