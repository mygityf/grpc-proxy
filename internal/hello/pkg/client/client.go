package main

import (
	"context"
	"google.golang.org/grpc"
	hellopb "grpc-proxy/api-gen/hello/v1"
	"grpc-proxy/internal/proxy/pkg/model"
	"log"
)

func main() {
	// conn to grpc proxy
	conn, err := grpc.Dial(model.GrpcProxyEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	response, err := hellopb.NewHelloServiceClient(conn).Hello(context.Background(),
		&hellopb.HelloRequest{
			Message: "hello",
		})
	if err != nil {
		log.Fatalf("Hello err:%v", err)
		return
	}
	log.Println("Hello response", response.Message)
}
