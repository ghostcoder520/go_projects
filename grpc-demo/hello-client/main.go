package main

import (
	"context"
	"fmt"
	pb "grpc-demo/hello-client/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到server端。此处禁用安全传输，没有验证和加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	defer conn.Close()

	// 创建sayhello服务的客户端，连接服务端
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用（这个SayHello方法具体在服务端实现），并返回结果
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "hhan"})

	fmt.Println(resp.GetResponseMsg())

}
