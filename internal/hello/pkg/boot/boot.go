package boot

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	hellopb "grpc-proxy/api-gen/hello/v1"
	"grpc-proxy/internal/hello/internal/api"
)

// register grpc service
func GrpcRegister(s *grpc.Server) {
	hellopb.RegisterHelloServiceServer(s, api.NewHelloServiceImpl())
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
}
