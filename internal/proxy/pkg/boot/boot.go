package boot

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	proxypb "grpc-proxy/api-gen/proxy/v1"
	"grpc-proxy/internal/proxy/internal/api"
)

// register grpc service
func GrpcRegister(s *grpc.Server) {
	proxypb.RegisterProxyRegisterServiceServer(s, api.NewProxyRegisterImpl())
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
}
