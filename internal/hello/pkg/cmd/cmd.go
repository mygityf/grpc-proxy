package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	hellopb "grpc-proxy/api-gen/hello/v1"
	proxypb "grpc-proxy/api-gen/proxy/v1"
	helloBoot "grpc-proxy/internal/hello/pkg/boot"
	"grpc-proxy/internal/proxy/pkg/model"
	"grpc-proxy/internal/worker/pkg/stop"
	"log"
	"net"
	"time"
)

const (
	grpcHelloServicePort = "9999"
)

// NewHelloCommand
func NewHelloCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello server",
		Short: "grpc hello server",
		Long:  "grpc hello server",
		Run: func(cmd *cobra.Command, args []string) {
			execute()
		},
	}
	return cmd
}

// start grpc proxy
func execute() {
	listenAddress := net.JoinHostPort("0.0.0.0", grpcHelloServicePort)
	// 启动外网代理端口
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Println(fmt.Sprintf("Listen outer at: %s", listenAddress))
	}
	// 试用 SSL 证书加密通道
	var grpcServer *grpc.Server
	grpcServer = grpc.NewServer()
	helloBoot.GrpcRegister(grpcServer)
	go grpcServer.Serve(listen)
	go func() {
		for {
			registerHelloServer()
			time.Sleep(time.Second * 10)
		}
	}()
	stop.SignalHandler()
}

func registerHelloServer() {
	// conn to grpc proxy
	conn, err := grpc.Dial(model.GrpcProxyEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()
	// export hello service name
	// var ExportHelloServiceName = _HelloService_serviceDesc.ServiceName
	registerRequest := &proxypb.RegisterRequest{
		ServiceName: hellopb.ExportHelloServiceName,
		Endpoint:    net.JoinHostPort("127.0.0.1", grpcHelloServicePort),
	}
	_, err = proxypb.NewProxyRegisterServiceClient(conn).Register(
		context.Background(), registerRequest)
	if err != nil {
		log.Fatalf("Register err: %v", err)
	}
}
