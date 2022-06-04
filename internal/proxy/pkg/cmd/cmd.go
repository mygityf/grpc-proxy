package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	proxyBoot "grpc-proxy/internal/proxy/pkg/boot"
	"grpc-proxy/internal/proxy/pkg/forward"
	"grpc-proxy/internal/proxy/pkg/proxy"
	"grpc-proxy/internal/worker/pkg/stop"
	"log"
	"net"
)

const (
	grpcProxyPort = "9991"
)

// NewGrpcProxyCommand
func NewGrpcProxyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grpc proxy",
		Short: "grpc proxy forward server",
		Long:  "grpc proxy forward server",
		Run: func(cmd *cobra.Command, args []string) {
			execute()
		},
	}
	return cmd
}

// start grpc proxy
func execute() {
	listenAddress := net.JoinHostPort("0.0.0.0", grpcProxyPort)
	// 启动外网代理端口
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Println(fmt.Sprintf("Listen at: %s", listenAddress))
	}
	// 试用 SSL 证书加密通道
	var grpcServer *grpc.Server
	grpcServer = grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(forward.Director)),
	)
	proxyBoot.GrpcRegister(grpcServer)
	go grpcServer.Serve(listen)
	stop.SignalHandler()
}
