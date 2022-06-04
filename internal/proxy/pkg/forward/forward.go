package forward

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-proxy/internal/proxy/internal/logic"
	"log"
	"strings"
)

// director 简单转发请求到新的网络地址上
func Director(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	serviceName := getServiceNameFromFullMethodName(fullMethodName)
	log.Println("proxy.forward.Director", fullMethodName, serviceName)
	conn, err := logic.GetRegisterImpl().GetConn(ctx, serviceName)
	return NewCtx(ctx), conn, err
}

// 获取gRPC服务名称
func getServiceNameFromFullMethodName(fullMethodName string) string {
	serviceNameList := strings.Split(fullMethodName, "/")
	if len(serviceNameList) <= 1 {
		return ""
	}
	return serviceNameList[1]
}

func NewCtx(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	// Copy the inbound metadata explicitly.
	outCtx, _ := context.WithCancel(ctx)
	outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
	return outCtx
}
